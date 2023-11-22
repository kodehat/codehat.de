package routes

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kodehat/codehat.de/pkg/types"
)

func Handle(ctx context.Context, host string, port int) {
	fs, err := fs.Sub(ctx.Value(types.ContextKeyStaticFiles).(embed.FS), "static")
	if err != nil {
		log.Fatal(err)
	}
	if ctx.Value(types.ContextKeyIsDebug).(bool) {
		dir, _ := os.Getwd()
		fs = os.DirFS(filepath.Join(dir, "static"))
	}
	http.Handle("/static/", staticHandler(ctx, http.StripPrefix("/static/", http.FileServer(http.FS(fs)))))
	http.Handle("/", dataMiddleware(ctx, http.HandlerFunc(serveRoot{}.handle)))
	http.Handle("/ip.html", dataMiddleware(ctx, http.HandlerFunc(serveIp{}.handle)))

	log.Printf("Listening on %s:%d\n", host, port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
