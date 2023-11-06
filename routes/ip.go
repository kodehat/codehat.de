package routes

import (
	"context"
	"net/http"
	"strings"
)

type serveIp struct {
	ctx context.Context
}

func (s serveIp) handle(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Host string
		IP   string
	}{Host: r.Host, IP: ReadUserIP(r)}
	renderPage(w, r, r.URL.Path, s.ctx, data)
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
