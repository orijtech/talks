package custom

import (
	"context"
	"log"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

type customExporter struct{}

func (ce *customExporter) ExportView(vd *view.Data) {
	log.Printf("Process me and send me to your metrics backend: %#v\n", vd)
}

var MRequestsIn = stats.Int64("demo/requests", "The number of requests received", "1")
var KeyMethod, _ = tag.NewKey("method")
var RequestsView = &view.View{Measure: MRequestsIn, Name: "demo/requests", Description: "Number of requests",
	Aggregation: view.Count(), TagKeys: []tag.Key{KeyMethod},
}
