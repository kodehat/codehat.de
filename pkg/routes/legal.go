package routes

import "net/http"

type serveLegalDetails struct {
}

func (s serveLegalDetails) handle(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, r.URL.Path)
}
