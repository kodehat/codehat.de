package routes

import (
	"fmt"
	"net/http"
)

const ROBOTS_TXT_CONTENT string = `User-agent: *
Disallow:
`

type serverRobotsTxt struct{}

func (s serverRobotsTxt) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, ROBOTS_TXT_CONTENT)
}
