        java.sql.Connection conn = new Connection(originalConn,
                                        // And passing this option to allow the spans
                                        // to be annotated with the SQL queries.
                                        // Please note that this could be a security concern
                                        // since it could reveal personally identifying information.
                                        Observability.OPTION_ANNOTATE_TRACES_WITH_SQL);
