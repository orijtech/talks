import io.opencensus.exporter.trace.jaeger.JaegerTraceExporter;
import io.opencensus.exporter.stats.prometheus.PrometheusStatsCollector;
import io.prometheus.client.exporter.HTTPServer;

public static void enableObservability() throws Exception {
    Observability.registerAllViews();

    // The trace exporter, for this demo we'll use Zipkin.
    JaegerTraceExporter.createAndRegister("http://127.0.0.1:14268/api/traces", "ocjdbc-demo");

    // The metrics exporter.
    PrometheusStatsCollector.createAndRegister();

    // Run the server as a daeon on address "localhost:8889".
    HTTPServer server = new HTTPServer("localhost", 8889);
}
