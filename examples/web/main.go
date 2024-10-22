package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
		UiVariant: "redoc",
	})

	w.Tag(web.Tag{
		Name:        "demo",
		Description: "demodescription",
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
		Method:      http.MethodGet,
		Path:        "/test/{name}",
		Description: "blabla",
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

	w.Api(web.Api{
		Method:    http.MethodGet,
		Path:      "/test",
		Parameter: web.Parameter{},
		Responses: web.Responses{
			StatusOK: ResponseTest{Say: "Hello world"},
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("content-type", "application/json")
			w.Write([]byte("test"))
		}),
	})

	w.Api(web.Api{
		Method:    http.MethodGet,
		Path:      "/sse",
		Parameter: web.Parameter{},
		Responses: web.Responses{
			StatusOK: ResponseTest{Say: "Hello world"},
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%v", r.Header)
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")

			for i := 0; i < 10; i++ {
				fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", i))
				time.Sleep(2 * time.Second)
				w.(http.Flusher).Flush()
			}

			<-r.Context().Done()
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
	log.Println("Visit http://localhost:8082/api/doc to view the documentation")
	if err := http.ListenAndServe(":8082", w.Server()); err != nil {
		panic(err)
	}
}
