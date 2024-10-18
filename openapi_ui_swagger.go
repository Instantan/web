package web

import "net/http"

type openapiUiSwagger struct {
	title  string
	docUrl string
}

func (or *openapiUiSwagger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		<meta name="description" content="SwaggerUI" />
		<title>` + or.title + `</title>
		<link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css" />
		<link rel="icon" href="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDI0IDI0IiBmaWxsPSJub25lIiBjb2xvcj0icmdiKDUyIDIxMSAxNTMpIiBzdHJva2U9ImN1cnJlbnRDb2xvciIgc3Ryb2tlLXdpZHRoPSIyIiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiIGNsYXNzPSJsdWNpZGUgbHVjaWRlLXdlYmhvb2sgaC02IHctNiBtci0yIHRleHQtZW1lcmFsZC00MDAiPjxwYXRoIGQ9Ik0xOCAxNi45OGgtNS45OWMtMS4xIDAtMS45NS45NC0yLjQ4IDEuOUE0IDQgMCAwIDEgMiAxN2MuMDEtLjcuMi0xLjQuNTctMiI+PC9wYXRoPjxwYXRoIGQ9Im02IDE3IDMuMTMtNS43OGMuNTMtLjk3LjEtMi4xOC0uNS0zLjFhNCA0IDAgMSAxIDYuODktNC4wNiI+PC9wYXRoPjxwYXRoIGQ9Im0xMiA2IDMuMTMgNS43M0MxNS42NiAxMi43IDE2LjkgMTMgMTggMTNhNCA0IDAgMCAxIDAgOCI+PC9wYXRoPjwvc3ZnPg==">

	</head>
	<body style="padding: 0px; margin: 0px;">
	<div id="swagger-ui"></div>
	<script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js" crossorigin></script>
	<script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-standalone-preset.js" crossorigin></script>
	<script>
		window.onload = () => {
		window.ui = SwaggerUIBundle({
			url: '` + or.docUrl + `',
			dom_id: '#swagger-ui',
			presets: [
			SwaggerUIBundle.presets.apis,
			SwaggerUIStandalonePreset
			],
			layout: "StandaloneLayout",
		});
		};
	</script>
	</body>
	</html>
	`))
}
