package main

import (
	"sync"
	"time"
)

func main() {
	spinSleep := func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			<-time.After(500 * time.Millisecond)
		}
	}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go spinSleep(&wg)
	}
	wg.Wait()
}

