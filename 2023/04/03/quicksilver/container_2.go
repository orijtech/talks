	defer func() {
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	}()

	var conn *grpc.ClientConn
	address := "localhost:" + resource.GetPort("9010/tcp")
	if err := pool.Retry(func() error {
		if conn != nil {
			return nil
		}
		var err error
		conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		return err
	}); err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}
