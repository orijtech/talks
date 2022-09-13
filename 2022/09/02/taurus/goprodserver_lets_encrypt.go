package main

import (
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello, TLS!"))
	})
	panic(http.Serve(autocert.NewListener("go.orijtech.com/taurus"), nil))
}
