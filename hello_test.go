package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if !strings.Contains(string(body), "Greetings") {
		t.Error("no contains greeting")
	}
}
