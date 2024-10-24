package generate

import (
	"bytes"

	"github.com/Instantan/web/internal/openapi"
)

const typescriptFetchClient = `

type ClientOptions = {
	url?: string
	beforeRequest?: (api: {
		method: any,
		path: any,
		params: {
			path: any,
			header: any,
			body: any,
			cookie: any,
			query: any
		}
	}) => void,
	afterRequest?: (response: {
		status: number,
		body: any
	}) => void
}

function createClient(options?: ClientOptions): Api {
	const url = options?.url ? options.url : ''
	return (async (api: any) => {
		if (options?.beforeRequest) {
			options.beforeRequest(api)
		}
		const query = new URLSearchParams()
		const queryObj = api?.params?.query || {}
		Object.keys(queryObj).forEach(name => query.set(name, queryObj[name]))
		const queryString = (query.size > 0 ? '?' + query.toString() : '')
		const pathParams = api?.params?.path || {}
		const path = api.path.replace(/{(\w+)}/g, (_, key) => 
			pathParams[key] !== undefined ? pathParams[key] : '{' + key + '}'
		)
		const resp = await fetch(url + path + queryString, {
			method: api.method,
			body: api?.params?.body,
		})
		const result = {
			status: resp.status,
			body: await resp.json()
		}
		if (options?.afterRequest) {
			options.afterRequest(result)
		}
		return result
	}) as Api
}

export {
	createClient
}
`

func GenerateTypescriptModels(api openapi.OpenAPI) []byte {
	t := newTsGenerator(bytes.NewBuffer([]byte{}))
	t.comment("Code generated by web DO NOT EDIT").newline()
	if api.Info.Title != "" {
		t.doc([]string{
			api.Info.Title + " " + api.Info.Version,
		})
	}
	if api.Paths.Len() > 0 {
		t.name("interface").s(" ").name("Api").s(" ").scope(func(t *tsGenerator) {
			for route, path := range api.Paths.Iterate() {
				for method, operation := range path.IterateOperations() {
					t.braces(func(t *tsGenerator) {
						t.name("api").colon().scope(func(t *tsGenerator) {
							t.name("method").colon().s("'" + method + "'").newline()
							t.name("path").colon().s("'" + route + "'").newline()

							querySchema := operation.QuerySchema()
							cookieSchema := operation.CookieSchema()
							pathSchema := operation.PathSchema()
							headerSchema := operation.HeaderSchema()

							t.name("params").colon().scope(func(t *tsGenerator) {
								if len(pathSchema.Properties) > 0 {
									t.name("path").colon().schema(pathSchema).newline()
								}
								if len(querySchema.Properties) > 0 {
									t.name("query").colon().schema(querySchema).newline()
								}
								if len(cookieSchema.Properties) > 0 {
									t.name("cookie").colon().schema(cookieSchema).newline()
								}
								if len(headerSchema.Properties) > 0 {
									t.name("header").colon().s("Record<string, string> & ").schema(headerSchema)
								} else {
									t.name("header?").colon().s("Record<string, string>")
								}
							})
							if operation.RequestBody != nil {
								if !operation.RequestBody.Required {
									t.name("body").colon()
								} else {
									t.name("body?").colon()
								}
								for _, content := range operation.RequestBody.Content {
									t.schema(content.Schema).union().newline()
								}
								t.marker()
							}
						})
					}).colon().name("Promise").generic(func(t *tsGenerator) {
						for code, response := range operation.Responses.Iterate() {
							t.scope(func(t *tsGenerator) {
								t.name("status").colon().s(code).newline()
								t.name("body").colon()
								for _, mediatype := range response.Content {
									t.schema(mediatype.Schema).union()
								}
								t.marker()
							}).union()
						}
						t.marker()
					}).semicolon().newline().newline()
				}
			}
			t.marker()
		}).newline().newline()
	}

	for name, schema := range api.Components.Schemas {
		t.name("type").s(" ").name(name).assign().schema(schema).newline()
	}

	t.s(typescriptFetchClient)

	return t.bytes()
}
