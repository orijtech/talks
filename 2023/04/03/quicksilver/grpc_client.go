import "google.golang.org/grpc"

func spinUpClient(targetAddr string) {
	conn, err := grpc.DialContext(ctx, targetAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Now you can use the gRPC connection in your generated gRPC service client.
}
