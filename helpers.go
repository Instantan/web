package web

import "net/http"

func Chain(middlewares ...Use) Use {
	return func(finalHandler http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			finalHandler = middlewares[i](finalHandler)
		}
		return finalHandler
	}
}

func mergeDefaults(defaults ...Defaults) Defaults {
	merged := Defaults{
		Query:     Query{},
		Header:    Header{},
		Cookie:    Cookie{},
		Body:      Body{},
		Responses: Responses{},
	}
	for _, def := range defaults {
		if len(def.Body.Description) > 0 {
			merged.Body.Description = def.Body.Description
		}
		if def.Body.Value != nil {
			merged.Body.Value = def.Body.Value
		}
		if len(def.Cookie) > 0 {
			for key, value := range def.Cookie {
				merged.Cookie[key] = value
			}
		}
		if len(def.Header) > 0 {
			for key, value := range def.Cookie {
				merged.Cookie[key] = value
			}
		}
		if len(def.Query) > 0 {
			for key, value := range def.Query {
				merged.Query[key] = value
			}
		}
		for status, value := range def.Responses.Iterate() {
			merged.Responses.Set(status, value)
		}
	}
	return merged
}
