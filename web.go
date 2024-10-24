package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Instantan/web/internal/generate"
	"github.com/Instantan/web/internal/openapi"
)

type Web struct {
	info                  Info
	contact               *Contact
	license               *License
	externalDocumentation ExternalDocumentation
	openapi               OpenApi
	typescriptApi         *TypescriptApi

	group Group
}

type Info struct {
	Title          string
	Version        string
	Summary        string
	Description    string
	TermsOfService string
}

type Contact struct {
	Name  string
	Url   string
	Email string
}

type License struct {
	Name       string
	Identifier string
}

type Tag struct {
	Name        string
	Description string
}

type ExternalDocumentation struct {
	Description string
	Url         string
}

type OpenApi struct {
	DocPath   string
	UiPath    string
	UiVariant string
}

type TypescriptApi struct {
	Path   string
	Writer io.Writer
}

func NewWeb() *Web {
	return &Web{
		group: Group{
			routes: &[]route{},
		},
	}
}

func (web *Web) Info(info Info) {
	assertIsNotEmpty("Info.Title", info.Title)
	assertIsNotEmpty("Info.Version", info.Version)
	web.info = info
}

func (web *Web) Contact(contact Contact) {
	web.contact = &contact
}

func (web *Web) License(license License) {
	web.license = &license
}

func (web *Web) Tag(tag Tag) {
	web.group.Tag(tag)
}

func (web *Web) ExternalDocumentation(externalDocumentation ExternalDocumentation) {
	assertIsNotEmpty("ExternalDocumentation.Url", externalDocumentation.Url)
	web.externalDocumentation = externalDocumentation
}

func (web *Web) OpenApi(openapi OpenApi) {
	assertIsNotEmpty("OpenApi.DocPath", openapi.DocPath)
	assertIsNotEmpty("OpenApi.UiPath", openapi.UiPath)
	assertIsNotEmpty("OpenApi.UiVariant", openapi.UiVariant)
	assertIsOneOf(openapi.UiVariant, []string{"scalar", "swagger", "redoc"})
	web.openapi = openapi
}

func (web *Web) Use(use Use) {
	web.group.Use(use)
}

func (web *Web) Api(route Api) {
	web.group.Api(route)
}

func (web *Web) Static(static Static) {
	web.group.Static(static)
}

func (web *Web) Group(group func(Group)) {
	web.group.Group(group)
}

func (web *Web) TypescriptApi(typescriptApi TypescriptApi) {
	web.typescriptApi = &typescriptApi
}

func (web *Web) Server() http.Handler {
	mux := http.NewServeMux()

	components := &openapi.Components{
		Schemas:         map[string]openapi.Schema{},
		Responses:       map[string]openapi.Response{},
		Parameters:      map[string]openapi.Parameter{},
		Examples:        map[string]openapi.Example{},
		RequestBodies:   map[string]openapi.RequestBody{},
		Headers:         map[string]openapi.Header{},
		SecuritySchemes: map[string]openapi.SecurityScheme{},
		Links:           map[string]openapi.Link{},
		Callbacks:       map[string]openapi.Callback{},
		PathItems:       map[string]openapi.PathItem{},
	}

	tags := &tags{
		tags:       &map[string]Tag{},
		references: []string{},
	}

	oa := openapi.OpenAPI{}
	oa.OpenApi = "3.1.0"
	oa.Info = web.info.openapiInfo()
	oa.Paths = *web.group.openapiPaths(mux, components, nil, tags)
	oa.Components = *components
	oa.Tags = tags.openapiTags()
	oa.Servers = []openapi.Server{}

	if web.contact != nil {
		oa.Info.Contact = web.contact.openapiContact()
	}
	if web.license != nil {
		oa.Info.License = web.license.openapiLicense()
	}

	schema, err := json.Marshal(oa)
	if err != nil {
		panic(err)
	}

	if web.openapi.DocPath != "" {
		mux.HandleFunc(http.MethodGet+" "+web.openapi.DocPath, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(200)
			w.Write(schema)
		})
	}
	if web.openapi.DocPath != "" && web.openapi.UiPath != "" {
		mux.Handle(http.MethodGet+" "+web.openapi.UiPath, web.openapi.httpHandler(web.info.Title))
	}

	if web.typescriptApi != nil {
		if web.typescriptApi.Writer == nil {
			web.typescriptApi.Writer, err = openOrCreateFile(web.typescriptApi.Path)
			if err != nil {
				panic(err)
			}
		}
		_, err := web.typescriptApi.Writer.Write(generate.GenerateTypescriptModels(oa))
		if err != nil {
			panic(err)
		}
	}

	return mux
}

func (info Info) openapiInfo() openapi.Info {
	return openapi.Info{
		Title:       info.Title,
		Summary:     info.Summary,
		Description: info.Description,
		Version:     info.Version,
	}
}

func (contact Contact) openapiContact() *openapi.Contact {
	return &openapi.Contact{
		Name:  contact.Name,
		Url:   contact.Url,
		Email: contact.Email,
	}
}

func (license License) openapiLicense() *openapi.License {
	return &openapi.License{
		Name:       license.Name,
		Identifier: license.Identifier,
	}
}

func (openapi OpenApi) httpHandler(title string) http.Handler {
	switch openapi.UiVariant {
	case "redoc":
		return &openapiUiRedoc{
			title:  title,
			docUrl: openapi.DocPath,
		}
	case "swagger":
		return &openapiUiSwagger{
			title:  title,
			docUrl: openapi.DocPath,
		}
	case "scalar":
		fallthrough
	default:
		return &openapiUiScalar{
			title:  title,
			docUrl: openapi.DocPath,
		}
	}
}
