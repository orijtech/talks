package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	cst := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello caller, your protocol is %s", r.Proto)
	}))
	defer cst.Close()
	res, err := cst.Client().Get(cst.URL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	hello, err := io.ReadAll(res.Body)
	if err != nil {	panic(err) }
	println(string(hello))
}
