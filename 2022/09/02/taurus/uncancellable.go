package main

import "time"

func dispatch(n int, fn func() error) chan error {
	ch := make(chan error, 1)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			if err := fn(); err != nil { ch <- err }
		}
	}()
	return ch
}

func main() {
	resCh := dispatch(1e7, func() error { time.Sleep(time.Second); return nil })
	if err := <-resCh; err != nil {
		panic(err)
	}
	println("Completed")
}
