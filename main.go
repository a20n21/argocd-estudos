package main

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello ARGOCD, foram feitas novas alterações nessa aplicação"))
    })

    http.Handle("/metrics", promhttp.Handler())

    http.ListenAndServe(":8080", nil)
}