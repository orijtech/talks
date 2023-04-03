package main

func testFetching(t *testing.T, sc *spanner.Client) {
	// Now your business logic in here as if you had a connection
	// to a production database.
}

func setupCloudSpannerEmulator(t *testing.T, fnToRun func(*testing.T, *spanner.Client)) {
	runOpts := &dockertest.RunOptions{
		Repository: "gcr.io/cloud-spanner-emulator/emulator",
		Tag:        "latest",
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}
	resource, err := pool.RunWithOptions(runOpts)
	if err != nil {
		t.Fatalf("Could not start resource: %s", err)
	}
	resource.Expire(30)
