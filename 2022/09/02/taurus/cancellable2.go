func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		defer cancel()
		select {
		case <-time.After(5 * time.Second):
		case <-sigCh:
		}
	}()

	resCh := cancellable(ctx, 1e7, func() error {
		time.Sleep(time.Second)
		return nil
	})
	if err := <-resCh; err != nil {
		panic(err)
	}
	println("Completed")
}
