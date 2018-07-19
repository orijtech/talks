package main

import (
	"log"
	"net/http"

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {
	h := &ochttp.Handler{Handler: http.HandlerFunc(search)}
	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	ctx, span := trace.StartSpan(r.Context(), "Search")
	defer span.End()

	// Use the context and the rest of the code goes below
	_ = ctx
}
