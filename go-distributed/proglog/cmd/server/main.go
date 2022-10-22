package main

import (
  "log"

  "github.com/dolpheyn/go-playground/go-distributed/proglog/internal/server"
)

func main() {
  server := server.NewHTTPServer(":8080")
  log.Fatal(server.ListenAndServe())
}
