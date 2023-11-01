package routes

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func renderPage(w http.ResponseWriter, r *http.Request, page string, ctx context.Context, data any) {
	page = filepath.Clean(page)
	page = strings.TrimPrefix(page, "/")

	functionsMap := template.FuncMap{
		"toLower": strings.ToLower,
		"toUpper": strings.ToUpper,
	}

	template := template.New("")
	var err error

	if ctx.Value("isDebug").(bool) {
		dir, _ := os.Getwd()
		template, err = template.Funcs(functionsMap).ParseGlob(filepath.Join(dir, "templates") + "/*.html")
	} else {
		template, err = template.Funcs(functionsMap).ParseFS(ctx.Value("templateFiles").(embed.FS), "templates/*.html")
	}

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	if template.Lookup(page) == nil {
		data := struct{}{}
		renderPage(w, r, "404.html", ctx, data)
		return
	}

	pageData := struct {
		Title string
		Data  any
	}{
		Data: data,
	}
	err = template.ExecuteTemplate(w, page, pageData)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

}
