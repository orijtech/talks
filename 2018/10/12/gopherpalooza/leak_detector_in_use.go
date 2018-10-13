package main

import (
	"bytes"
	"fmt"
	"runtime"
	"time"
)

func doWork() {
	for {
		<-time.After(1 * time.Second)
	}
	_ = 1 + 1
}

func main() {
	for i := 0; i < 10; i++ {
		go doWork()
	}
	leakDetector()
}

func parseStackSections(section []byte) {
	// Fill in the function to parse goroutine information
}

func leakDetector() {
	buf := make([]byte, 2048)
	for {
		n := runtime.Stack(buf, true)
		splits := bytes.Split(buf[:n], []byte("\n\n"))
		for i, split := range splits {
			fmt.Printf("#%d:\n%s\n", i, split)
		}
		<-time.After(2 * time.Second)
		// Reuse the buffer for next refresh
		copy(buf[len(buf)-n:n], bytes.Repeat([]byte{0x00}, n))
	}
}
