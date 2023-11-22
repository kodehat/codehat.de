package main

import (
	"context"
	"embed"
	"flag"
	"log"
	"os"

	"github.com/kodehat/codehat.de/pkg/routes"
	"github.com/kodehat/codehat.de/pkg/types"
)

// build flags
var (
	BuildTime  string = "N/A"
	CommitHash string = "N/A"
	GoVersion  string = "N/A"
	GitTag     string = "N/A"
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
	var isVersion bool
	flag.BoolVar(&isVersion, "version", false, "can be used to print build information like version and timestamp")
	flag.Parse()
	if port <= 0 {
		log.Fatal("Port must be greater than zero!")
		os.Exit(1)
	}
	buildInformation := types.BuildInformation{
		BuildTime: BuildTime, CommitHash: CommitHash, GoVersion: GoVersion, GitTag: GitTag,
	}
	if isVersion {
		log.Println("Build time: ", BuildTime)
		log.Println("Commit: ", CommitHash)
		log.Println("Go version: ", GoVersion)
		log.Println("Git tag: ", GitTag)
		os.Exit(0)
	}
	if isDebug {
		log.Println("Debug mode is turned on.")
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, types.ContextKeyBuildInformation, buildInformation)
	routesCtx := context.WithValue(ctx, types.ContextKeyTemplatesFiles, templateFiles)
	routesCtx = context.WithValue(routesCtx, types.ContextKeyStaticFiles, staticFiles)
	routesCtx = context.WithValue(routesCtx, types.ContextKeyIsDebug, isDebug)
	routes.Handle(routesCtx, host, port)
}
