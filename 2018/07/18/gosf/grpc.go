package main

import (
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"google.golang.org/grpc"
)

func main() {
	_ = view.Register(ocgrpc.DefaultServerViews...)
	_ = view.Register(ocgrpc.DefaultClientViews...)
	_ = grpc.NewServer(grpc.StatsHandler(new(ocgrpc.ServerHandler)))
}
