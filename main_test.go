package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	PingHandler(w, req)

	expected := "pong\n" // Fprintln adiciona \n

	if w.Body.String() != expected {
		t.Errorf("esperado %s, veio %s", expected, w.Body.String())
	}
}

func TestRootHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	RootHandler(w, req)

	expected := "Hello ARGOCD + Jaeger + OpenTelemetry 🚀\n"

	if w.Body.String() != expected {
		t.Errorf("esperado %s, veio %s", expected, w.Body.String())
	}
}