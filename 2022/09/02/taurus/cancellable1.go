package main

import (
	"context"
	"os"
	"os/signal"
	"time"
)

func dispatch(ctx context.Context, n int, fn func() error) chan error {
	ch := make(chan error, 1)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			if err := ctx.Err(); err != nil {
				if err != context.Canceled { ch <- err }
				break
			}

			if err := fn(); err != nil { ch <- err }
		}
	}()
	return ch
}

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

	resCh := dispatch(ctx, 1e7, func() error {
		time.Sleep(time.Second)
		return nil
	})
	if err := <-resCh; err != nil {
		panic(err)
	}
	println("Completed")
}
