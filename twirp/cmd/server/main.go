package main

import (
    "net/http"

    "github.com/example/internal/haberdasherserver"
    "github.com/example/rpc/haberdasher"
)

func main() {
  server := &haberdasherserver.Server{} // implements Haberdasher interface
  twirpHandler := haberdasher.NewHaberdasherServer(server)

  http.ListenAndServe(":8080", twirpHandler)
}