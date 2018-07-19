import (
	"go.opencensus.io/stats"
	"go.opencensus.io/view"
)

var LatencyView = &view.View{
	Measure: MLatency, Name: "demo/latency", Description: "Latency per call",
	TagKeys:     []tag.Key{KeyMethod},
	Aggregation: view.Distribution(0, 5, 10, 20, 50, 100, 200, 300, 500, 800),
}
var RequestsView = &view.View{Measure: MRequestsIn, Name: "demo/requests",
	Description: "Number of requests", Aggregation: view.Count(),
}

func init() {
	if err := view.Register(LatencyView, RequestsView); err != nil {
		log.Fatalf("Failed to register custom views: %v", err)
	}
}
