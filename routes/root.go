package routes

import (
	"context"
	"net/http"
)

type serveRoot struct {
	ctx context.Context
}

func (s serveRoot) handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.URL.Path = "index.html"
	}

	data := struct{}{}
	renderPage(w, r, r.URL.Path, s.ctx, data)
}
