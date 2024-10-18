package openapi

import (
	"encoding/json"
	"errors"
	"net/http"
)

// This is the root object of the OpenAPI document.
type OpenAPI struct {
	// REQUIRED. This string MUST be the version number of the OpenAPI Specification
	// that the OpenAPI document uses. The openapi field SHOULD be used by
	// tooling to interpret the OpenAPI document. This is not related to the
	// API info.version string.
	OpenApi string `json:"openapi"` // `json:"age,omitempty"`
	// REQUIRED. Provides metadata about the API. The metadata MAY be used by
	// tooling as required.
	Info Info `json:"info"`
	// The default value for the $schema keyword within Schema Objects
	// contained within this OAS document. This MUST be in the form of a URI.
	JsonSchemaDialect string `json:"jsonSchemaDialect,omitempty"`
	// An array of Server Objects, which provide connectivity information to a
	// target server. If the servers property is not provided, or is an empty array,
	// the default value would be a Server Object with a url value of /.
	Servers []Server `json:"servers"`
	// The available paths and operations for the API.
	Paths Paths `json:"paths,omitempty"`
	// The incoming webhooks that MAY be received as part of this API and that
	// the API consumer MAY choose to implement. Closely related to the callbacks
	// feature, this section describes requests initiated other than by an API call,
	// for example by an out of band registration. The key name is a unique string to
	// refer to each webhook, while the (optionally referenced) Path Item Object
	// describes a request that may be initiated by the API provider and the expected
	// responses. An example is available.
	Webhooks map[string]PathItem/*Reference*/ `json:"webooks,omitempty"`
	// An element to hold various schemas for the document.
	Components Components `json:"components,omitempty"`
	// A declaration of which security mechanisms can be used across the API. The list
	// of values includes alternative security requirement objects that can be used.
	// Only one of the security requirement objects need to be satisfied to authorize
	// a request. Individual operations can override this definition. To make security
	//optional, an empty security requirement ({}) can be included in the array.
	Security SecurityRequirement `json:"security,omitempty"`
	// A list of tags used by the document with additional metadata. The order of the
	// tags can be used to reflect on their order by the parsing tools. Not all tags
	// that are used by the Operation Object must be declared. The tags that are not
	// declared MAY be organized randomly or based on the tools’ logic. Each tag name
	// in the list MUST be unique.
	Tags []Tag `json:"tags,omitempty"`
	// Additional external documentation.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}

type Info struct {
	// REQUIRED. The title of the API.
	Title string `json:"title"`
	// A short summary of the API.
	Summary string `json:"summary,omitempty"`
	// A description of the API. [CommonMark] syntax MAY be used for rich text
	// representation.
	Description string `json:"description,omitempty"`
	// A URL to the Terms of Service for the API. This MUST be in the form of a URL.
	TermsOfService string `json:"termsOfService,omitempty"`
	// The contact information for the exposed API.
	Contact *Contact `json:"contact,omitempty"`
	// The license information for the exposed API.
	License *License `json:"license,omitempty"`
	// REQUIRED. The version of the OpenAPI document (which is distinct from the
	// OpenAPI Specification version or the API implementation version).
	Version string `json:"version"`
}

type Contact struct {
	// The identifying name of the contact person/organization.
	Name string `json:"name,omitempty"`
	// The URL pointing to the contact information. This MUST be in the form of a URL
	Url string `json:"url,omitempty"`
	// The email address of the contact person/organization. This MUST be in the form
	// of an email address.
	Email string `json:"email,omitempty"`
}

type License struct {
	// REQUIRED. The license name used for the API.
	Name string `json:"name,omitempty"`
	// An [SPDX-Licenses] expression for the API. The identifier field is mutually
	// exclusive of the url field.
	Identifier string `json:"identifier,omitempty"`
	// A URL to the license used for the API. This MUST be in the form of a URL.
	// The url field is mutually exclusive of the identifier field.
	Url string `json:"url,omitempty"`
}

type Server struct {
	// REQUIRED. A URL to the target host. This URL supports Server Variables and
	// MAY be relative, to indicate that the host location is relative to the location
	// where the OpenAPI document is being served. Variable substitutions will be made
	// when a variable is named in {brackets}.
	Url string `json:"url,omitempty"`
	// An optional string describing the host designated by the URL. [CommonMark]
	// syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// A map between a variable name and its value. The value is used for substitution
	// in the server’s URL template.
	Variables map[string]ServerVariable `json:"variables,omitempty"`
}

type ServerVariable struct {
	// An enumeration of string values to be used if the substitution options are from a
	// limited set. The array MUST NOT be empty.
	Enum []string `json:"enum,omitempty"`
	// REQUIRED. The default value to use for substitution, which SHALL be sent if
	// an alternate value is not supplied. Note this behavior is different than the
	// Schema Object’s treatment of default values, because in those cases parameter
	// values are optional. If the enum is defined, the value MUST exist in the enum’s
	// values.
	Default string `json:""`
	// An optional description for the server variable. [CommonMark] syntax MAY be used
	// for rich text representation.
	Description string `json:"description,omitempty"`
}

// A relative path to an individual endpoint. The field name MUST begin with a
// forward slash (/). The path is appended (no relative URL resolution) to the
// expanded URL from the Server Object’s url field in order to construct the full
// URL. Path templating is allowed. When matching URLs, concrete (non-templated)
// paths would be matched before their templated counterparts. Templated paths with
// the same hierarchy but different templated names MUST NOT exist as they are
// identical. In case of ambiguous matching, it’s up to the tooling to decide which one to use.
type Paths = OrderedMap[string, PathItem]

type PathItem struct {
	// Allows for a referenced definition of this path item. The referenced
	// structure MUST be in the form of a Path Item Object. In case a Path Item
	// Object field appears both in the defined object and the referenced object,
	// the behavior is undefined. See the rules for resolving Relative References.
	Ref string `json:"$ref,omitempty"`
	// An optional, string summary, intended to apply to all operations in this path.
	Summary string `json:"summary,omitempty"`
	// An optional, string description, intended to apply to all operations in this path.
	// [CommonMark] syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// A definition of a GET operation on this path.
	Get *Operation `json:"get,omitempty"`
	// A definition of a PUT operation on this path.
	Put *Operation `json:"put,omitempty"`
	// A definition of a POST operation on this path.
	Post *Operation `json:"post,omitempty"`
	// A definition of a DELETE operation on this path.
	Delete *Operation `json:"delete,omitempty"`
	// A definition of a OPTIONS operation on this path.
	Options *Operation `json:"options,omitempty"`
	// A definition of a HEAD operation on this path.
	Head *Operation `json:"head,omitempty"`
	// A definition of a PATCH operation on this path.
	Patch *Operation `json:"patch,omitempty"`
	// A definition of a TRACE operation on this path.
	Trace *Operation `json:"trace,omitempty"`
	// An alternative server array to service all operations in this path.
	Servers []Server `json:"servers,omitempty"`
	// A list of parameters that are applicable for all the operations described
	// under this path. These parameters can be overridden at the operation level,
	// but cannot be removed there. The list MUST NOT include duplicated parameters.
	// A unique parameter is defined by a combination of a name and location. The list
	// can use the Reference Object to link to parameters that are defined at the
	// OpenAPI Object’s components/parameters.
	Parameters []Parameter/*Reference*/ `json:"parameters,omitempty"`
}

type Operation struct {
	// A list of tags for API documentation control. Tags can be used for logical
	// grouping of operations by resources or any other qualifier.
	Tags []string `json:"tags,omitempty"`
	// A short summary of what the operation does.
	Summary string `json:"summary,omitempty"`
	// A verbose explanation of the operation behavior. [CommonMark] syntax MAY
	// be used for rich text representation.
	Description string `json:"description,omitempty"`
	// Additional external documentation for this operation.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
	// Unique string used to identify the operation. The id MUST be unique among
	// all operations described in the API. The operationId value is case-sensitive.
	// Tools and libraries MAY use the operationId to uniquely identify an operation,
	// therefore, it is RECOMMENDED to follow common programming naming conventions.
	OperationId string `json:"operationId,omitempty"`
	// A list of parameters that are applicable for this operation. If a parameter is
	// already defined at the Path Item, the new definition will override it but can
	// never remove it. The list MUST NOT include duplicated parameters. A unique parameter
	// is defined by a combination of a name and location. The list can use the Reference
	// Object to link to parameters that are defined at the OpenAPI Object’s components/parameters.
	Parameters []Parameter/*Reference*/ `json:"parameters"`
	// The request body applicable for this operation. The requestBody is fully supported in HTTP
	// methods where the HTTP 1.1 specification [RFC7231] Section 4.3.1 has explicitly defined
	// semantics for request bodies. In other cases where the HTTP spec is vague
	// (such as GET, HEAD and DELETE), requestBody is permitted but does not have well-defined
	// semantics and SHOULD be avoided if possible.
	RequestBody *RequestBody/*Reference*/ `json:"requestBody,omitempty"`
	// The list of possible responses as they are returned from executing this operation.
	Responses Responses `json:"responses,omitempty"`
	// A map of possible out-of band callbacks related to the parent operation. The key is a
	// unique identifier for the Callback Object. Each value in the map is a Callback Object
	// that describes a request that may be initiated by the API provider and the expected
	// responses.
	Callbacks map[string]Callback/*Reference*/ `json:"callbacks,omitempty"`
	// Declares this operation to be deprecated. Consumers SHOULD refrain from usage of the
	// declared operation. Default value is false.
	Deprecated bool `json:"deprecated"`
	// A declaration of which security mechanisms can be used for this operation. The list
	// of values includes alternative security requirement objects that can be used. Only
	// one of the security requirement objects need to be satisfied to authorize a request.
	// To make security optional, an empty security requirement ({}) can be included in the
	// array. This definition overrides any declared top-level security. To remove a top-level
	// security declaration, an empty array can be used.
	Security SecurityRequirement `json:"security"`
	// An alternative server array to service this operation. If an alternative server object
	// is specified at the Path Item Object or Root level, it will be overridden by this value.
	Servers []Server `json:"server"`
}

type RequestBody struct {
	// A brief description of the request body. This could contain examples of use.
	// [CommonMark] syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// REQUIRED. The content of the request body. The key is a media type or media type
	// range, see [RFC7231] Appendix D, and the value describes it. For requests that match
	// multiple keys, only the most specific key is applicable. e.g. text/plain overrides
	// text/*
	Content map[string]MediaType `json:"content"`
	// Determines if the request body is required in the request. Defaults to false.
	Required bool `json:"required"`
}

type MediaType struct {
	// The schema defining the content of the request, response, or parameter.
	Schema Schema `json:"schema"`
	// Example of the media type. The example object SHOULD be in the correct format as
	// specified by the media type. The example field is mutually exclusive of the examples
	// field. Furthermore, if referencing a schema which contains an example, the example
	// value SHALL override the example provided by the schema.
	Example any `json:"example,omitempty"`
	// Examples of the media type. Each example object SHOULD match the media type and
	// specified schema if present. The examples field is mutually exclusive of the example
	// field. Furthermore, if referencing a schema which contains an example, the examples
	// value SHALL override the example provided by the schema.
	Examples map[string]Example/*Reference*/ `json:"examples,omitempty"`
	// A map between a property name and its encoding information. The key, being the property
	// name, MUST exist in the schema as a property. The encoding object SHALL only apply to
	// requestBody objects when the media type is multipart or application/x-www-form-urlencoded.
	Encoding map[string]Encoding `json:"encoding"`
}

type Encoding struct {
	// The Content-Type for encoding a specific property. Default value depends on the property
	// type: for object - application/json; for array – the default is defined based on the
	// inner type; for all other cases the default is application/octet-stream. The value can
	// be a specific media type (e.g. application/json), a wildcard media type (e.g. image/*),
	// or a comma-separated list of the two types.
	ContentType string `json:"contentType"`
	// A map allowing additional information to be provided as headers, for example
	// Content-Disposition. Content-Type is described separately and SHALL be ignored in
	// this section. This property SHALL be ignored if the request body media type is not
	// a multipart.
	Headers map[string]Header/*Reference*/ `json:"headers,omitempty"`
	// Describes how a specific property value will be serialized depending on its type.
	// See Parameter Object for details on the style property. The behavior follows the same
	// values as query parameters, including default values. This property SHALL be ignored
	// if the request body media type is not application/x-www-form-urlencoded or multipart/form-data.
	// If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
	Style string `json:"style"`
	// When this is true, property values of type array or object generate separate parameters
	// for each value of the array, or key-value-pair of the map. For other types of properties
	// this property has no effect. When style is form, the default value is true. For all other
	// styles, the default value is false. This property SHALL be ignored if the request body
	// media type is not application/x-www-form-urlencoded or multipart/form-data. If a value
	// is explicitly defined, then the value of contentType (implicit or explicit) SHALL be
	// ignored.
	Explode bool `json:"explode"`
	// Determines whether the parameter value SHOULD allow reserved characters, as defined
	// by [RFC3986] Section 2.2 :/?#[]@!$&'()*+,;= to be included without percent-encoding.
	// The default value is false. This property SHALL be ignored if the request body media
	// type is not application/x-www-form-urlencoded or multipart/form-data. If a value is
	// explicitly defined, then the value of contentType (implicit or explicit) SHALL be
	// ignored.
	AllowReserved bool
}

type Example struct {
	// Short description for the example.
	Summary string `json:"summary,omitempty"`
	// Long description for the example. [CommonMark] syntax MAY be used
	// for rich text representation.
	Description string `json:"description,omitempty"`
	// Embedded literal example. The value field and externalValue field are mutually exclusive.
	// To represent examples of media types that cannot naturally represented in JSON or YAML,
	// use a string value to contain the example, escaping where necessary.
	Value any `json:"value,omitempty"`
	// A URI that points to the literal example. This provides the capability to reference examples
	// that cannot easily be included in JSON or YAML documents. The value field and externalValue
	// field are mutually exclusive. See the rules for resolving Relative References.
	ExternalValue string `json:"externalValue,omitempty"`
}

type Responses struct {
	// The documentation of responses other than the ones declared for specific HTTP response
	// codes. Use this field to cover undeclared responses.
	Default Response/*Reference*/ `json:"default,omitempty"`
	// Any HTTP status code can be used as the property name, but only one property per code,
	//to describe the expected response for that HTTP status code. This field MUST be enclosed
	// in quotation marks (for example, “200”) for compatibility between JSON and YAML.
	// To define a range of response codes, this field MAY contain the uppercase wildcard
	// character X. For example, 2XX represents all response codes between [200-299]. Only
	// the following range definitions are allowed: 1XX, 2XX, 3XX, 4XX, and 5XX. If a
	// response is defined using an explicit code, the explicit code definition takes
	// precedence over the range definition for that code.
	HTTPStatusCodeResponses map[string]Response /*Reference*/
}

type Response struct {
	// REQUIRED. A description of the response. [CommonMark] syntax MAY be used for rich
	// text representation.
	Description string `json:"description"`
	// Maps a header name to its definition. [RFC7230] Page 22 states header names are case
	// insensitive. If a response header is defined with the name "Content-Type", it SHALL
	// be ignored.
	Headers map[string]Header/*Reference*/ `json:"headers"`
	// A map containing descriptions of potential response payloads. The key is a media type
	// or media type range, see [RFC7231] Appendix D, and the value describes it. For responses
	// that match multiple keys, only the most specific key is applicable. e.g. text/plain
	// overrides text/*
	Content map[string]MediaType `json:"content"`
	// A map of operations links that can be followed from the response. The key of the map
	// is a short name for the link, following the naming constraints of the names for Component
	// Objects.
	Links map[string]Link /*Reference*/
}

type Link struct {
	// A relative or absolute URI reference to an OAS operation. This field is mutually
	// exclusive of the operationId field, and MUST point to an Operation Object. Relative
	// operationRef values MAY be used to locate an existing Operation Object in the OpenAPI
	// definition. See the rules for resolving Relative References.
	OperationRef string `json:"operationRef,omitempty"`
	// The name of an existing, resolvable OAS operation, as defined with a unique operationId.
	// This field is mutually exclusive of the operationRef field.
	OperationId string `json:"operationId,omitempty"`
	// A map representing parameters to pass to an operation as specified with operationId or
	// identified via operationRef. The key is the parameter name to be used, whereas the value
	// can be a constant or an expression to be evaluated and passed to the linked operation.
	// The parameter name can be qualified using the parameter location [{in}.]{name} for
	// operations that use the same parameter name in different locations (e.g. path.id).
	Parameters map[string]any `json:"parameters,omitempty"`
	// A literal value or {expression} to use as a request body when calling the target
	// operation.
	RequestBody any `json:"requestBody,omitempty"`
	// A description of the link. [CommonMark] syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// A server object to be used by the target operation.
	Server Server `json:"server,omitempty"`
}

type Header struct {
	// A brief description of the parameter. This could contain examples of use. [CommonMark]
	// syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// Determines whether this parameter is mandatory. If the parameter location is "path",
	// this property is REQUIRED and its value MUST be true. Otherwise, the property MAY
	// be included and its default value is false.
	Required bool `json:"required"`
	// Specifies that a parameter is deprecated and SHOULD be transitioned out of usage.
	// Default value is false.
	Deprecated bool `json:"deprecated"`
	// Sets the ability to pass empty-valued parameters. This is valid only for query
	// parameters and allows sending a parameter with an empty value. Default value is
	// false. If style is used, and if behavior is n/a (cannot be serialized), the value
	// of allowEmptyValue SHALL be ignored. Use of this property is NOT RECOMMENDED, as it
	// is likely to be removed in a later revision.
	AllowEmptyValue bool `json:"allowEmptyValue"`
	// Describes how the parameter value will be serialized depending on the type of the parameter value. Default values
	// (based on value of in): for query - form; for path - simple; for header - simple;
	// for cookie - form.
	Style string `json:"style,omitempty"`
	// When this is true, parameter values of type array or object generate separate parameters
	// for each value of the array or key-value pair of the map. For other types of parameters
	// this property has no effect. When style is form, the default value is true. For all
	// other styles, the default value is false.
	Explode bool `json:"explode"`
	// Determines whether the parameter value SHOULD allow reserved characters, as defined by
	// [RFC3986] Section 2.2 :/?#[]@!$&'()*+,;= to be included without percent-encoding. This
	// property only applies to parameters with an in value of query. The default value is
	// false.
	AllowReserved bool `json:"allowReserved"`
	// The schema defining the type used for the parameter.
	Schema Schema `json:"schema,omitempty"`
	// Example of the parameter’s potential value. The example SHOULD match the specified
	// schema and encoding properties if present. The example field is mutually exclusive of
	// the examples field. Furthermore, if referencing a schema that contains an example, the
	// example value SHALL override the example provided by the schema. To represent examples
	// of media types that cannot naturally be represented in JSON or YAML, a string value can
	// contain the example with escaping where necessary.
	Example any `json:"example,omitempty"`
	// Examples of the parameter’s potential value. Each example SHOULD contain a value in the
	// correct format as specified in the parameter encoding. The examples field is mutually
	// exclusive of the example field. Furthermore, if referencing a schema that contains an
	// example, the examples value SHALL override the example provided by the schema.
	Examples map[string]Example/*Reference*/ `json:"examples,omitempty"`
	// A map containing the representations for the parameter. The key is the media type and
	// the value describes it. The map MUST only contain one entry.
	Content map[string]MediaType `json:"content,omitempty"`
}

// A Path Item Object, or a reference to one, used to define a callback request and
// expected responses. A complete example is available.
type Callback map[string]PathItem /*Reference*/

type Parameter struct {
	// REQUIRED. The name of the parameter. Parameter names are case sensitive.
	// If in is "path", the name field MUST correspond to a template expression
	// occurring within the path field in the Paths Object. See Path Templating
	// for further information.
	// If in is "header" and the name field is "Accept", "Content-Type" or "Authorization",
	// the parameter definition SHALL be ignored.
	// For all other cases, the name corresponds to the parameter name used by the in property.
	Name string `json:"name"`
	// REQUIRED. The location of the parameter. Possible values are "query", "header",
	// "path" or "cookie".
	In string `json:"in"`
	// A brief description of the parameter. This could contain examples of use. [CommonMark]
	// syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// Determines whether this parameter is mandatory. If the parameter location is "path",
	// this property is REQUIRED and its value MUST be true. Otherwise, the property MAY
	// be included and its default value is false.
	Required bool `json:"required"`
	// Specifies that a parameter is deprecated and SHOULD be transitioned out of usage.
	// Default value is false.
	Deprecated bool `json:"deprecated"`
	// Sets the ability to pass empty-valued parameters. This is valid only for query
	// parameters and allows sending a parameter with an empty value. Default value is
	// false. If style is used, and if behavior is n/a (cannot be serialized), the value
	// of allowEmptyValue SHALL be ignored. Use of this property is NOT RECOMMENDED, as it
	// is likely to be removed in a later revision.
	AllowEmptyValue bool `json:"allowEmptyValue"`
	// Describes how the parameter value will be serialized depending on the type of the parameter value. Default values
	// (based on value of in): for query - form; for path - simple; for header - simple;
	// for cookie - form.
	Style string `json:"style,omitempty"`
	// When this is true, parameter values of type array or object generate separate parameters
	// for each value of the array or key-value pair of the map. For other types of parameters
	// this property has no effect. When style is form, the default value is true. For all
	// other styles, the default value is false.
	Explode bool `json:"explode"`
	// Determines whether the parameter value SHOULD allow reserved characters, as defined by
	// [RFC3986] Section 2.2 :/?#[]@!$&'()*+,;= to be included without percent-encoding. This
	// property only applies to parameters with an in value of query. The default value is
	// false.
	AllowReserved bool `json:"allowReserved"`
	// The schema defining the type used for the parameter.
	Schema Schema `json:"schema,omitempty"`
	// Example of the parameter’s potential value. The example SHOULD match the specified
	// schema and encoding properties if present. The example field is mutually exclusive of
	// the examples field. Furthermore, if referencing a schema that contains an example, the
	// example value SHALL override the example provided by the schema. To represent examples
	// of media types that cannot naturally be represented in JSON or YAML, a string value can
	// contain the example with escaping where necessary.
	Example any `json:"example,omitempty"`
	// Examples of the parameter’s potential value. Each example SHOULD contain a value in the
	// correct format as specified in the parameter encoding. The examples field is mutually
	// exclusive of the example field. Furthermore, if referencing a schema that contains an
	// example, the examples value SHALL override the example provided by the schema.
	Examples map[string]Example/*Reference*/ `json:"examples,omitempty"`
	// A map containing the representations for the parameter. The key is the media type and
	// the value describes it. The map MUST only contain one entry.
	Content map[string]MediaType `json:"content,omitempty"`
}

type Reference struct {
	// REQUIRED. The reference identifier. This MUST be in the form of a URI.
	Ref string `json:"$ref"`
	// A short summary which by default SHOULD override that of the referenced component.
	// If the referenced object-type does not allow a summary field, then this field has no effect.
	Summary string `json:"summary,omitempty"`
	// A description which by default SHOULD override that of the referenced component.
	// [CommonMark] syntax MAY be used for rich text representation. If the referenced
	// object-type does not allow a description field, then this field has no effect.
	Description string `json:"description,omitempty"`
}

type Components struct {
	// An object to hold reusable Schema Objects.
	Schemas map[string]Schema `json:"schemas,omitempty"`
	// An object to hold reusable Response Objects.
	Responses map[string]Response/*Reference*/ `json:"responses,omitempty"`
	// An object to hold reusable Parameter Objects.
	Parameters map[string]Parameter/*Reference*/ `json:"parameters,omitempty"`
	// An object to hold reusable Parameter Objects.
	Examples map[string]Example/*Reference*/ `json:"examples,omitempty"`
	// An object to hold reusable Request Body Objects.
	RequestBodies map[string]RequestBody/*Reference*/ `json:"requestBodies,omitempty"`
	// An object to hold reusable Header Objects.
	Headers map[string]Header/*Reference*/ `json:"headers,omitempty"`
	// An object to hold reusable Header Objects.
	SecuritySchemes map[string]SecurityScheme/*Reference*/ `json:"securitySchemes,omitempty"`
	// An object to hold reusable Link Objects.
	Links map[string]Link/*Reference*/ `json:"links,omitempty"`
	// An object to hold reusable Callback Objects.
	Callbacks map[string]Callback/*Reference*/ `json:"callbacks,omitempty"`
	// An object to hold reusable Path Item Object.
	PathItems map[string]PathItem/*Reference*/ `json:"pathItems,omitempty"`
}

type SecurityScheme struct {
	// REQUIRED. The type of the security scheme. Valid values are "apiKey", "http",
	// "mutualTLS", "oauth2", "openIdConnect".
	Type string `json:"type"`
	// A description for security scheme. [CommonMark] syntax MAY be used for rich
	// text representation.
	Description string `json:"description,omitempty"`
	// REQUIRED. The name of the header, query or cookie parameter to be used.
	Name string `json:"name"`
	// REQUIRED. The location of the API key. Valid values are "query", "header" or "cookie".
	In string `json:"in"`
	// REQUIRED. The name of the HTTP Authorization scheme to be used in the Authorization
	// header as defined in [RFC7235] Section 5.1. The values used SHOULD be registered
	// in the IANA Authentication Scheme registry.
	Scheme string `json:"scheme"`
	// A hint to the client to identify how the bearer token is formatted. Bearer tokens are
	// usually generated by an authorization server, so this information is primarily for
	// documentation purposes.
	BearerFormat string `json:"bearerFormat,omitempty"`
	// REQUIRED. An object containing configuration information for the flow types supported.
	Flows OAuthFlows `json:"flows"`
	// REQUIRED. OpenId Connect URL to discover OAuth2 configuration values. This MUST be in
	// the form of a URL. The OpenID Connect standard requires the use of TLS.
	OpenIdConnectUrl string `json:"openIdConnect"`
}

type OAuthFlows struct {
	// Configuration for the OAuth Implicit flow
	Implicit OAuthFlow `json:"implicit,omitempty"`
	// Configuration for the OAuth Resource Owner Password flow
	Password OAuthFlow `json:"password,omitempty"`
	// Configuration for the OAuth Client Credentials flow. Previously called application in OpenAPI 2.0.
	ClientCredentials OAuthFlow `json:"clientCredentials,omitempty"`
	// Configuration for the OAuth Authorization Code flow. Previously called accessCode in OpenAPI 2.0.
	AuthorizationCode OAuthFlow `json:"authorizationCode,omitempty"`
}

type OAuthFlow struct {
	// REQUIRED. The authorization URL to be used for this flow. This MUST be in
	// the form of a URL. The OAuth2 standard requires the use of TLS.
	AuthorizationUrl string `json:"authorizationUrl"`
	// REQUIRED. The token URL to be used for this flow. This MUST be in the form of a URL.
	// The OAuth2 standard requires the use of TLS.
	TokenUrl string `json:"tokenUrl"`
	// The URL to be used for obtaining refresh tokens. This MUST be in the form of a URL.
	// The OAuth2 standard requires the use of TLS.
	RefreshUrl string `json:"refreshUrl,omitempty"`
	// REQUIRED. The available scopes for the OAuth2 security scheme. A map between the scope
	// name and a short description for it. The map MAY be empty.
	Scopes map[string]string `json:"scopes"`
}

// Each name MUST correspond to a security scheme which is declared in the Security
// Schemes under the Components Object. If the security scheme is of type "oauth2" or
// "openIdConnect", then the value is a list of scope names required for the execution,
// and the list MAY be empty if authorization does not require a specified scope. For
// other security scheme types, the array MAY contain a list of role names which are
// required for the execution, but are not otherwise defined or exchanged in-band.
type SecurityRequirement map[string][]string

type Tag struct {
	// REQUIRED. The name of the tag.
	Name string `json:"name,omitempty"`
	// A description for the tag. [CommonMark] syntax MAY be used for rich text
	// representation.
	Description string `json:"description,omitempty"`
	// Additional external documentation for this tag.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}

type ExternalDocumentation struct {
	// A description for the tag. [CommonMark] syntax MAY be used for rich text
	// representation.
	Description string `json:"description,omitempty"`
	// REQUIRED. The URL for the target documentation. This MUST be in the form of a URL.
	Url string `json:"url"`
}

func (u Responses) MarshalJSON() ([]byte, error) {
	m := u.HTTPStatusCodeResponses
	if m == nil {
		m = map[string]Response /*Reference*/ {}
	}
	if u.Default.Content != nil {
		m["default"] = u.Default
	}
	return json.Marshal(m)
}

func (u OpenAPI) MarshalJSON() ([]byte, error) {
	if err := errors.Join(
		requireOpenAPIField("Info.Title", u.Info.Title),
		requireOpenAPIField("Info.Version", u.Info.Version),
	); err != nil {
		return nil, err
	}
	type Alias OpenAPI
	return json.Marshal(Alias(u))
}

func (m Schema) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		n := map[string]any{}
		n["$ref"] = m.Ref
		return json.Marshal(n)
	}
	type Alias Schema
	return json.Marshal(Alias(m))
}

func (p PathItem) IterateOperations() func(func(string, *Operation) bool) {
	return func(yield func(status string, value *Operation) bool) {
		if p.Get != nil {
			if !yield(http.MethodGet, p.Get) {
				return
			}
		}
		if p.Head != nil {
			if !yield(http.MethodHead, p.Head) {
				return
			}
		}
		if p.Options != nil {
			if !yield(http.MethodOptions, p.Options) {
				return
			}
		}
		if p.Patch != nil {
			if !yield(http.MethodPatch, p.Patch) {
				return
			}
		}
		if p.Post != nil {
			if !yield(http.MethodPatch, p.Post) {
				return
			}
		}
		if p.Delete != nil {
			if !yield(http.MethodDelete, p.Delete) {
				return
			}
		}
		if p.Put != nil {
			if !yield(http.MethodPut, p.Put) {
				return
			}
		}
		if p.Trace != nil {
			if !yield(http.MethodTrace, p.Trace) {
				return
			}
		}
	}
}

func (r Responses) Iterate() func(func(string, Response) bool) {
	return func(yield func(status string, response Response) bool) {
		if r.Default.Content != nil {
			if !yield("0", r.Default) {
				return
			}
		}
		for status, response := range r.HTTPStatusCodeResponses {
			if !yield(status, response) {
				return
			}
		}
	}
}
