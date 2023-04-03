	// Now you can run the code that required a gRPC connection
	// to Cloud Spanner, as you would in production.
	opts := []option.ClientOption{option.WithGRPCConn(conn)}
	shutdownCtx, shutdown := context.WithCancel(context.Background())
	project, instance, db := "project", "instance", "test-db"
	err := InitSpannerDB(shutdownCtx, project, instance, db, stmts, opts)
	require.Nil(t, err)

	fullDBName := fmt.Sprintf("projects/%s/instances/%s/databases/%s", project, instance, db)

	// Boom, now create the Cloud Spanner Client.
	sc, err := spanner.NewClient(shutdownCtx, fullDBName, opts...)
	require.Nil(t, err)
	defer func() {
		sc.Close()
		shutdown()
	}()

	fnToRun(t, sc)
}
