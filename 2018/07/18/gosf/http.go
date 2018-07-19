package main

import (
	"net/http"

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
)

func main() {
	// Start of metrics/stats enabling by registering views
	_ = view.Register(ochttp.DefaultServerViews...)
	_ = view.Register(ochttp.DefaultClientViews...)

	// Start of trace propagation configuration for HTTP
	// For a trace enabled HTTP client/transport
	_ = &http.Client{Transport: &ochttp.Transport{Base: http.DefaultTransport}}
	// For a trace enabled HTTP server/handler
	_ = &ochttp.Handler{Handler: http.DefaultServeMux}
}
