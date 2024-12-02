package main

import (
  "fmt"
  "os"
)

func main() {
  defer fmt.Println("Execution finished")

  filename := "customers.txt"
  file, err := os.Open(filename)

  defer file.Close()

  if err != nil {
    panic("The indicated file was not found or is damaged")
  } 

  data := make([]byte, 100)
  count, err := file.Read(data)

  if err != nil {
    panic("Could not read the indicated file")
  }

  fmt.Printf("Read %d bytes:\n", count)
  fmt.Println(string(data[:count]))
} 

