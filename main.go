package main

import (
  "fmt"
  "net"
  "net/http"
)

func main() {
  http.HandleFunc("/ip", func(w http.ResponseWriter, r *http.Request) {
    ip, _, _ := net.SplitHostPort(r.RemoteAddr)
    fmt.Fprintf(w,ip + "\n\n")
  })
  http.ListenAndServe(":8080", nil)
}
