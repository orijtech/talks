package main

func setupServer() {
	ln, err := net.Listen("tcp", ":0") // To auto-assign you an available local port
	if err != nil {
		panic(err)
	}

	gsrv := grpc.NewServer()
	// Bind the gRPC service definitions being tested to this server
	gsrv.RegisterService(svcToTest, nil)

	go func() {
		defer ln.Close()
		if err := gsrv.Listen(ln); err != nil {
			panic(err)
		}
	}()

	// Now you can use the ln.Addr().String() to pass into the gRPC client
	targetAddrForClient := ln.Addr().String()

        // Continue down below..
}
