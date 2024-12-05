package main

import (
  "fmt"
  "net/http"
  "encoding/json"
)

type Payload struct {
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"` 
}
  
func handleGreetings(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.WriteHeader(http.StatusMethodNotAllowed) 
    return
  }

  p := &Payload{}

  decoder := json.NewDecoder(r.Body)
  defer r.Body.Close()

  err := decoder.Decode(&p) 
  if err != nil { 
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  } 
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Hello %s %s", p.FirstName, p.LastName) 
  }
  
func main() {
  
  http.HandleFunc("/greetings", handleGreetings)

  fmt.Println("Server is running and listening on 8080...")
  
  http.ListenAndServe(":8080", nil)
  
} 
