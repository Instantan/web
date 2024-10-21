<img src="https://github.com/Instantan/web/blob/6ab9f9f33b150fb04f4b3ad2f91ae3f455b2baf1/website/src/assets/social-preview.png" alt="Web: The Zero-Dependency Go Web Framework" />

> Warning: This package is currently (not even) in alpha, not really tested yet and should not be used yet

## Todos
- Find a good api for defining content types
- Find a good api for defining sockets
- Add redoc

## Key Features
- **Zero Dependencies**: Built entirely on Go's standard library. No external packages required.
- **OpenAPI Integration**: Automatically generate OpenAPI specifications for your APIs, enhancing documentation and interoperability.
- **TypeScript API generator**: Automatically generate TypeScript definitions for your Go APIs, ensuring type safety across your full-stack application.

## Quick Start

Install with:
```bash
go get github.com/Instantan/web
```

Example usage:
```go
package main

import (
  "log"
  "net/http"
  "github.com/Instantan/web"
)

func main() {
  w := web.NewWeb()

  w.Info(web.Info{
    Title: "MyProject",
    Version: "0.0.1",
  })

  w.OpenApi(web.OpenApi{
    DocPath:   "/api/doc.json",
    UiPath:    "/api/doc",
    UiVariant: "scalar",
  })

  w.Api(web.Api{
    Method: http.MethodGet,
    Path:   "/hello/{name}",
    Parameter: web.Parameter{
      Path: web.Path{
        "name": web.PathParam{
          Description: "The name to say hello to",
          Value:       "world",
        },
      },
    },
    Responses: web.Responses{
      StatusOK: "Hello World",
    },
    Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte("Hello " + r.PathValue("name")))
    }),
  })

  log.Println("Server listening on :8080")
  log.Println("Visit http://localhost:8080/api/doc to view the documentation")
  if err := http.ListenAndServe(":8080", w.Server()); err != nil {
    panic(err)
  }
}
```

---

[![Go Report Card](https://goreportcard.com/badge/github.com/Instantan/web)](https://goreportcard.com/report/github.com/Instantan/web)

