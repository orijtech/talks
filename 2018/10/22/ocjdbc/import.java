import io.opencensus.integration.ocjdbc.Connection;
import io.opencensus.integration.ocjdbc.Observability;

public static void main(String ...args) {
    // After you've enabled your connection, we'll now wrap it
    java.sql.Connection conn = new Connection(originalConn);

    // Continue using your connection like you did in your normal programs
}

public static void enableObservability() {
    // Enable observability with OpenCensus
    Observability.registerAllViews();

    // Now create the trace and metrics exporters
}
