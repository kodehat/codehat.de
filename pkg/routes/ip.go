package routes

import (
	"net/http"
	"strings"

	"github.com/kodehat/codehat.de/pkg/types"
)

type serveIp struct {
}

func (s serveIp) handle(w http.ResponseWriter, r *http.Request) {
	ip := ReadUserIP(r)
	data := struct {
		IP     string
		IsIPv4 bool
	}{IP: ip, IsIPv4: IsIPv4(ip)}
	routeData := r.Context().Value(types.ContextKeyRouteData).(*types.RouteData)
	(*routeData).Local = data
	renderPage(w, r, r.URL.Path)
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	if lastColonIdx := strings.LastIndex(IPAddress, ":"); lastColonIdx > -1 {
		IPAddress = IPAddress[:lastColonIdx]
	}
	return IPAddress
}

func IsIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

func IsIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}
