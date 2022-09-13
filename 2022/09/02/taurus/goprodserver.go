package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
            fmt.Fprintf(rw, "The time now is: %s:🚨😎\n", time.Now())
	})
	panic(http.ListenAndServe(":8888", nil))
}
