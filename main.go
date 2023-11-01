package main

import (
	"context"
	"embed"
	"flag"
	"log"
	"os"

	"github.com/kodehat/codehat.de/routes"
)

//go:embed templates
var templateFiles embed.FS

//go:embed static
var staticFiles embed.FS

func main() {
	var host string
	flag.StringVar(&host, "host", "0.0.0.0", "host to listen on")
	var port int
	flag.IntVar(&port, "port", 1313, "port to listen on")
	var isDebug bool
	flag.BoolVar(&isDebug, "debug", false, "can be used to load local templates")
	flag.Parse()
	if port <= 0 {
		log.Fatal("Port must be greater than zero!")
		os.Exit(1)
	}
	ctx := context.Background()
	routesCtx := context.WithValue(ctx, "templateFiles", templateFiles)
	routesCtx = context.WithValue(routesCtx, "staticFiles", staticFiles)
	routesCtx = context.WithValue(routesCtx, "isDebug", isDebug)
	routes.Handle(routesCtx, host, port)
}
