package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
)

type alwaysSameRoundTripper int
func (asrt *alwaysSameRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	res := &http.Response{
		Status: "200 OK",StatusCode: 200, ProtoMajor: 1, ProtoMinor: 1,
		Header: http.Header{"x-backend": {"foo-bar"}}, Body: ioutil.NopCloser(strings.NewReader("Never left the network!")),
	}
	return res, nil
}
func main() {
	client := &http.Client{Transport: new(alwaysSameRoundTripper)}
	res, err := client.Get("https://example.org/expensive/")
	if err != nil {	panic(err) }
	resBlob, err := httputil.DumpResponse(res, true)
	if err != nil { panic(err) }
	println(string(resBlob))
}
