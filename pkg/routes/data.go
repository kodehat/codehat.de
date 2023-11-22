package routes

import (
	"context"
	"net/http"
	"strings"

	"github.com/kodehat/codehat.de/pkg/types"
)

func dataMiddleware(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		firstDotIdx := strings.Index(host, ".")
		domain := "localhost.localdomain"
		if firstDotIdx > -1 {
			domain = host[(firstDotIdx + 1):]
		}

		globalRouteData := types.GlobalRouteData{
			BuildInformation: ctx.Value(types.ContextKeyBuildInformation).(types.BuildInformation),
			Host:             host,
			Domain:           domain,
			IsDebug:          ctx.Value(types.ContextKeyIsDebug).(bool),
		}
		routeData := types.RouteData{
			Global: globalRouteData,
		}

		routeCtx := context.WithValue(ctx, types.ContextKeyRouteData, &routeData)
		next.ServeHTTP(w, r.WithContext(routeCtx))
	})
}
