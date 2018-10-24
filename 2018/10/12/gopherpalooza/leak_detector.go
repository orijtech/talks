package main

import (
	"bytes"
	"fmt"
	"runtime"
	"time"
)

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
	}
}
