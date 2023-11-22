package types

import "net/http"

type GlobalRouteData struct {
	BuildInformation
	Host    string
	Domain  string
	IsDebug bool
}

type RouteData struct {
	Global GlobalRouteData
	Local  any
}

type RouteHandler interface {
	handle(w http.ResponseWriter, r *http.Request)
}
