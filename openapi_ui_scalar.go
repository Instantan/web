package web

import "net/http"

type openapiUiScalar struct {
	title  string
	docUrl string
}

func (or *openapiUiScalar) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<!doctype html>
	<html>
	<head>
		<title>` + or.title + `</title>
		<meta charset="utf-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		<link rel="icon" href="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDI0IDI0IiBmaWxsPSJub25lIiBjb2xvcj0icmdiKDUyIDIxMSAxNTMpIiBzdHJva2U9ImN1cnJlbnRDb2xvciIgc3Ryb2tlLXdpZHRoPSIyIiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiIGNsYXNzPSJsdWNpZGUgbHVjaWRlLXdlYmhvb2sgaC02IHctNiBtci0yIHRleHQtZW1lcmFsZC00MDAiPjxwYXRoIGQ9Ik0xOCAxNi45OGgtNS45OWMtMS4xIDAtMS45NS45NC0yLjQ4IDEuOUE0IDQgMCAwIDEgMiAxN2MuMDEtLjcuMi0xLjQuNTctMiI+PC9wYXRoPjxwYXRoIGQ9Im02IDE3IDMuMTMtNS43OGMuNTMtLjk3LjEtMi4xOC0uNS0zLjFhNCA0IDAgMSAxIDYuODktNC4wNiI+PC9wYXRoPjxwYXRoIGQ9Im0xMiA2IDMuMTMgNS43M0MxNS42NiAxMi43IDE2LjkgMTMgMTggMTNhNCA0IDAgMCAxIDAgOCI+PC9wYXRoPjwvc3ZnPg==">

	</head>
	<body>
		<script
		id="api-reference"
		data-url="` + or.docUrl + `"></script>
		<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
	</body>
	</html>
	`))
}
