package main

import (
    "fmt"
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(200); w.Write([]byte("ok"))
    })
    http.HandleFunc("/api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, `{"message":"hello"}`)
    })
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":8080", nil)
}
