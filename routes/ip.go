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
	ip := ReadUserIP(r)
	data := struct {
		Host   string
		IP     string
		IsIPv4 bool
	}{Host: r.Host, IP: ip, IsIPv4: IsIPv4(ip)}
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

func IsIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

func IsIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}
