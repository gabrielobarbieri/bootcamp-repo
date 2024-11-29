package main

import (
  "fmt"
  "errors"
)

type Product struct {
  ID int
  Price float64
  Name string
  Description string
  Category string
}

var Products []Product

func (p Product) Save() {
  Products = append(Products, p)
}

func (p Product) GetAll() {
  for _, product := range Products {
    fmt.Println(product)
  }
}

func getById(id int) (Product, error)  {
  for _, product := range Products {
    if product.ID == id {
      return product, nil
    }
  }
  return Product{}, errors.New("No product with this id found!")
}

func main() {
  p1 := Product{
    ID: 1,
    Price: 100.0,
    Name: "Carro",
    Description: "Fast",
    Category: "Classificados",
  }
  p2 := Product{
    ID: 2,
    Price: 50.0,
    Name: "Roupa",
    Description: "Bonita",
    Category: "Moda",
  }

  fmt.Printf("Array products contém %d elementos\n", len(Products))

  p1.Save()

  fmt.Printf("Array products contém %d elementos\n", len(Products))

  p2.Save()

  fmt.Printf("Array products contém %d elementos\n", len(Products))

  p2.GetAll()  

  p, err := getById(2)

  if err != nil {
    return
  }

  fmt.Println(p)
}
