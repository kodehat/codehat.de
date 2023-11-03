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
)

func Handle(ctx context.Context, host string, port int) {
	fs, err := fs.Sub(ctx.Value("staticFiles").(embed.FS), "static")
	if err != nil {
		log.Fatal(err)
	}
	if ctx.Value("isDebug").(bool) {
		dir, _ := os.Getwd()
		fs = os.DirFS(filepath.Join(dir, "static"))
	}
	http.Handle("/static/", staticHandler(ctx, http.StripPrefix("/static/", http.FileServer(http.FS(fs)))))

	http.HandleFunc("/", serveRoot{ctx: ctx}.handle)
	http.HandleFunc("/ip.html", serveIp{ctx: ctx}.handle)

	log.Printf("Listening on %s:%d\n", host, port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
