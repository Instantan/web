package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Instantan/web"
)

type ResponseTest struct {
	Say string `json:"say"`
}

func main() {
	w := web.NewWeb()

	w.Info(web.Info{
		Title:   "Bla",
		Version: "0.0.1",
	})

	w.Contact(web.Contact{
		Name: "",
	})

	w.OpenApi(web.OpenApi{
		DocPath:   "/api/doc.json",
		UiPath:    "/api/doc",
		UiVariant: "scalar",
	})

	w.Group(func(g web.Group) {
		g.Defaults(web.Defaults{
			Query: web.Query{
				"": web.QueryParam{},
			},
		})
	})

	w.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf(r.RemoteAddr)
			next.ServeHTTP(w, r)
		})
	})

	w.Group(func(w web.Group) {
		w.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				log.Println("in group")
				next.ServeHTTP(w, r)
			})
		})
	})

	w.Api(web.Api{
		Method: http.MethodGet,
		Path:   "/test/{name}",
		Parameter: web.Parameter{
			Path: web.Path{
				"name": web.PathParam{
					Description: "The name to say hello to",
					Value:       "world", // example value
				},
			},
		},
		Responses: web.Responses{
			StatusOK: ResponseTest{Say: "Hello world"},
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("content-type", "application/json")
			res, _ := json.Marshal(struct {
				Say string `json:"name"`
			}{
				Say: "Hello " + r.PathValue("name"),
			})
			w.Write(res)
		}),
	})

	w.Static(web.Static{
		PathPrefix: "/",
		SpaMode:    true,
		FS:         http.Dir("./static"),
	})

	w.TypescriptApi(web.TypescriptApi{
		Path: "api.ts",
	})

	log.Println("Server listening on :8082")
	log.Println("Visist http://localhost:8082/api/doc to view the documentation")
	if err := http.ListenAndServe(":8082", w.Server()); err != nil {
		panic(err)
	}
}
