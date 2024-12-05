package main

import (
  "fmt"
  "net/http"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
  // response header -> response body 
  if r.Method == http.MethodGet {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "pong")
  } else {
      w.WriteHeader(http.StatusMethodNotAllowed)
      fmt.Fprintf(w, "Method now allowed")
    }
}

func main() {
  // Define endpoint -> Create route and router handler -> implement router handler logic -> start the server
  http.HandleFunc("/ping", handlePing)

  fmt.Println("Server is listening on :8080...")

  http.ListenAndServe(":8080", nil)
}
