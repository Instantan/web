package web

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Instantan/web/internal/openapi"
)

type Use func(next http.Handler) http.Handler

type Api struct {
	Method      string
	Path        string
	Tags        []string
	OperationId string
	Summary     string
	Description string
	Parameter   Parameter
	Responses   Responses
	Handler     http.Handler
}

type Group struct {
	routes *[]route
}

type Static struct {
	PathPrefix  string
	Tags        []string
	Summary     string
	Description string
	FS          http.FileSystem
	SpaMode     bool
}

type route struct {
	use      *Use
	api      *Api
	static   *Static
	defaults *Defaults
	group    *Group
}

func (g Group) Use(use Use) {
	*g.routes = append(*g.routes, route{
		use: &use,
	})
}

func (g Group) Api(api Api) {
	assertIsOneOf(strings.ToUpper(api.Method), []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodTrace,
	})
	assertIsNotEmpty("Get.Path", api.Path)
	assertIsNotNil("Get.Handler", api.Handler)
	*g.routes = append(*g.routes, route{
		api: &api,
	})
}

func (g Group) Static(static Static) {
	*g.routes = append(*g.routes, route{
		static: &static,
	})
}

func (g Group) Defaults(defaults Defaults) {
	*g.routes = append(*g.routes, route{
		defaults: &defaults,
	})
}

func (g Group) Group(group func(Group)) {
	gr := Group{
		routes: &[]route{},
	}
	group(gr)
	*g.routes = append(*g.routes, route{
		group: &gr,
	})
}

func (g Group) openapiPaths(mux *http.ServeMux, components *openapi.Components, use Use, defaults Defaults) *openapi.Paths {
	paths := &openapi.Paths{}

	for i := range *g.routes {
		r := (*g.routes)[i]
		if r.api != nil {
			api := r.api

			p, _ := paths.Get(api.Path)

			operation := &openapi.Operation{
				OperationId: api.OperationId,
				Tags:        api.Tags,
				Summary:     api.Summary,
				Description: api.Description,
				Responses: openapi.Responses{
					HTTPStatusCodeResponses: map[string]openapi.Response{},
				},
			}

			if api.Parameter.Body.Value != nil {
				operation.RequestBody.Required = !api.Parameter.Body.Optional
				operation.RequestBody.Description = api.Parameter.Body.Description
				content := operation.RequestBody.Content["text/*"]
				content.Example = api.Parameter.Body.Value
				s := *openapi.ValueToSchema(api.Parameter.Body.Value)
				if s.TypeName == "" {
					content.Schema = s
				} else {
					content.Schema.Ref = "#/components/schemas/" + s.TypeName
					components.Schemas[s.TypeName] = s
				}

				operation.RequestBody.Content["application/json"] = content
			}
			if len(api.Parameter.Cookie) > 0 {
				for key, value := range api.Parameter.Cookie {
					operation.Parameters = append(operation.Parameters, openapi.Parameter{
						Name:        key,
						In:          "cookie",
						Description: value.Description,
						Required:    !value.Optional,
						Schema:      *openapi.ValueToSchema(value.Value),
						Example:     value.Value,
					})
				}
			}
			if len(api.Parameter.Path) > 0 {
				for key, value := range api.Parameter.Path {
					operation.Parameters = append(operation.Parameters, openapi.Parameter{
						Name:        key,
						In:          "path",
						Description: value.Description,
						Required:    true,
						Schema:      *openapi.ValueToSchema(value.Value),
						Example:     value.Value,
					})
				}
			}
			if len(api.Parameter.Query) > 0 {
				for key, value := range api.Parameter.Query {
					operation.Parameters = append(operation.Parameters, openapi.Parameter{
						Name:        key,
						In:          "query",
						Description: value.Description,
						Required:    !value.Optional,
						Schema:      *openapi.ValueToSchema(value.Value),
						Example:     value.Value,
					})
				}
			}
			if len(api.Parameter.Header) > 0 {
				for key, value := range api.Parameter.Header {
					operation.Parameters = append(operation.Parameters, openapi.Parameter{
						Name:        key,
						In:          "header",
						Description: value.Description,
						Required:    !value.Optional,
						Schema:      *openapi.ValueToSchema(value.Value),
						Example:     value.Value,
					})
				}
			}

			createResponse := func(description string, value any) openapi.Response {
				content := openapi.MediaType{}
				s := *openapi.ValueToSchema(value)
				content.Example = value
				if s.TypeName == "" {
					content.Schema = s
				} else {
					content.Schema.Ref = "#/components/schemas/" + s.TypeName
					components.Schemas[s.TypeName] = s
				}

				return openapi.Response{
					Description: description,
					Content: map[string]openapi.MediaType{
						"application/json": content,
					},
				}
			}

			for status, value := range api.Responses.Iterate() {
				if status == 0 {
					operation.Responses.Default = createResponse("Default", value)
					continue
				}
				operation.Responses.HTTPStatusCodeResponses[strconv.Itoa(status)] = createResponse(http.StatusText(status), value)
			}

			switch api.Method {
			case http.MethodGet:
				p.Get = operation
			case http.MethodPut:
				p.Put = operation
			case http.MethodPost:
				p.Post = operation
			case http.MethodDelete:
				p.Delete = operation
			case http.MethodHead:
				p.Head = operation
			case http.MethodPatch:
				p.Patch = operation
			case http.MethodTrace:
				p.Trace = operation
			case http.MethodConnect:
				p.Trace = operation
			case http.MethodOptions:
				p.Options = operation
			}

			paths.Set(api.Path, p)
			if use == nil {
				mux.Handle(api.Method+" "+api.Path, api.Handler)
			} else {
				mux.Handle(api.Method+" "+api.Path, use(api.Handler))
			}
		} else if r.group != nil {
			group := r.group
			for path, item := range group.openapiPaths(mux, components, use, defaults).Iterate() {
				paths.Set(path, item)
			}
		} else if r.defaults != nil {
			defaults = mergeDefaults(*r.defaults)
		} else if r.use != nil {
			if use != nil {
				use = Chain(use, *r.use)
			} else {
				use = *r.use
			}
		} else if r.static != nil {
			var handler http.Handler
			if use != nil {
				handler = http.StripPrefix(r.static.PathPrefix, use(http.FileServer(r.static.FS)))
			} else {
				handler = http.StripPrefix(r.static.PathPrefix, http.FileServer(r.static.FS))
			}
			if r.static.SpaMode {
				mux.Handle(http.MethodGet+" "+r.static.PathPrefix, createSpaModeRedirect(handler))
			} else {
				mux.Handle(http.MethodGet+" "+r.static.PathPrefix, handler)
			}
		}
	}
	return paths
}
