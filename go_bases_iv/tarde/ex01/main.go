package main

import (
  "fmt"
  "os"
)

func main() {
  defer fmt.Println("Execution finished")

  filename := "customers.txt"
  _, err := os.Open(filename)

  if err != nil {
    panic("The indicated file was not found or is damaged")
  } 
} 
