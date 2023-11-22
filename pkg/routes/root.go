package routes

import (
	"net/http"
)

type serveRoot struct {
}

func (s serveRoot) handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.URL.Path = "index.html"
	}
	renderPage(w, r, r.URL.Path)
}
