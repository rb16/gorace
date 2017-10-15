package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleStructAdd(t *testing.T) {
	r := request(t, "/?first=13&second=32")
	rw := httptest.NewRecorder()

	handleStructAdd(rw, r)
	if rw.Code == 500 {
		t.Fatal("internal server error:" + rw.Body.String())
	}
}

func BenchmarkHandleStructAdd(b *testing.B) {
	r := request(b, "/?first=20&second=30")
	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		handleStructAdd(rw, r)
	}
}

func request(t testing.TB, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}
