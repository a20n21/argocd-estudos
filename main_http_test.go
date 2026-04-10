package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Body.String() != "pong" {
		t.Errorf("esperado pong, veio %s", w.Body.String())
	}
}