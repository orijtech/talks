package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int, wgg *sync.WaitGroup) {
			defer wgg.Done()
			_ = 10
			<-time.After(100 * time.Millisecond)
		}(i, &wg)
	}
	wg.Wait()
}
