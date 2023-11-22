package routes

import (
	"embed"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/kodehat/codehat.de/pkg/types"
)

func renderPage(w http.ResponseWriter, r *http.Request, page string) {
	page = filepath.Clean(page)
	page = strings.TrimPrefix(page, "/")

	functionsMap := template.FuncMap{
		"toLower": strings.ToLower,
		"toUpper": strings.ToUpper,
	}

	template := template.New("")
	var err error

	if r.Context().Value(types.ContextKeyIsDebug).(bool) {
		dir, _ := os.Getwd()
		template, err = template.Funcs(functionsMap).ParseGlob(filepath.Join(dir, "templates") + "/*.html")
	} else {
		template, err = template.Funcs(functionsMap).ParseFS(r.Context().Value(types.ContextKeyTemplatesFiles).(embed.FS), "templates/*.html")
	}

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	if template.Lookup(page) == nil {
		renderPage(w, r, "404.html")
		return
	}

	pageData := struct {
		Title string
		Data  types.RouteData
	}{
		Data: *(r.Context().Value(types.ContextKeyRouteData).(*types.RouteData)),
	}
	err = template.ExecuteTemplate(w, page, pageData)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
