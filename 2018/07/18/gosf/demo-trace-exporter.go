package main

import (
	"context"
	"fmt"
	"time"

	"go.opencensus.io/trace"
)

type customPrintExporter int

func (ce *customPrintExporter) ExportSpan(sd *trace.SpanData) {
    fmt.Printf("Name: %s\nTraceID: %x\nSpanID: %x\nParentSpanID: %x\nStartTime: %s\nEndTime: %s\nAnnotations: %+v\n",
		sd.Name, sd.TraceID, sd.SpanID, sd.ParentSpanID, sd.StartTime, sd.EndTime, sd.Annotations)
}

func main() {
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(0.999)})
	trace.RegisterExporter(new(customPrintExporter))
	_, span := trace.StartSpan(context.Background(), "sample")
	span.Annotate([]trace.Attribute{trace.Int64Attribute("invocations", 1)}, "Invoked it")
	span.End()
	<-time.After(500 * time.Millisecond)
}
