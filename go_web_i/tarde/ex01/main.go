package main

import (
  "fmt"
  "encoding/json"
  "os"
)

type Product struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Quantity int `json:"quantity"`
  CodeValue string `json:"codeValue"`
  IsPublished bool `json:"isPublished"`
  Expiration string `json:"expiration"`
  Price float64 `json:"price"`
} 

func main() {
  // Create a slice of product -> Decode the json file into the slice

  file, err := os.Open("./products.json")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()

  products := []Product{}

  decoder := json.NewDecoder(file)
  err = decoder.Decode(&products)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(products)
}
