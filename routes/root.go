package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/kodehat/codehat.de/pkg/types"
)

type serveRoot struct {
	ctx context.Context
}

func (s serveRoot) handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.URL.Path = "index.html"
	}
	log.Printf("Host is %s\n", r.Host)

	data := struct {
		BuildInformation types.BuildInformation
		Host             string
	}{BuildInformation: s.ctx.Value("buildInformation").(types.BuildInformation),
		Host: r.Host}
	renderPage(w, r, r.URL.Path, s.ctx, data)
}
