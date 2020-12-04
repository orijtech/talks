package main

import _ "net/http/pprof"

func main() {
	go func() {
		log.Println(http.ListenAndServe(":3338", nil))
	}()
	// Rest of your logic goes here
}
