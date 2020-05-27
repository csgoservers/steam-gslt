package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

// testRoundTrip is a type to use as test Transport
type testRoundTrip func(req *http.Request) *http.Response

// RoundTrip implementation to call our fake function
func (t testRoundTrip) RoundTrip(req *http.Request) (*http.Response, error) {
	return t(req), nil
}

// newTestClient creates our fake client
func newTestClient(fn testRoundTrip) *http.Client {
	return &http.Client{Transport: testRoundTrip(fn)}
}

func TestServiceUrl(t *testing.T) {
	client := newTestClient(func(req *http.Request) *http.Response {
		if "http://localhost/hello/v1?key=abc" != req.URL.String() {
			t.Errorf("Service URL is not valid: %s", req.URL.String())
		}
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       ioutil.NopCloser(bytes.NewBufferString("")),
		}
	})
	service := &gameServerService{client, "http://localhost", "abc"}
	service.get("hello")
}
