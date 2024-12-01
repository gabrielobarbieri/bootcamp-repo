package main

import (
  "fmt"
  "errors"
)

type Product interface {
  Price() float64
}

type smallProduct struct {
  productPrice float64
  productType string 
}

type mediumProduct struct {
  productPrice float64
  productType string
}

type bigProduct struct {
  productPrice float64
  productType string 
}

// criar func factory -> métodos -> criar interface product (com metodo price)
func (s smallProduct) Price() float64 {
  return s.productPrice 
}

func (m mediumProduct) Price() float64 {
  return m.productPrice + ((m.productPrice * 0.03) * 2)
}

func (b bigProduct) Price() float64 {
  return b.productPrice + (b.productPrice * 0.06) + 2500
}

// This func should receive a struct of type Product
func factory(productPrice float64, productType string) (Product, error) {
  switch productType {
    case "small":
      return smallProduct{productPrice, productType}, nil
    case "medium":
      return mediumProduct{productPrice, productType}, nil
    case "big":
      return bigProduct{productPrice, productType}, nil
    default:
      return nil, errors.New("No product with this type")
  }
}

func main() {
  bigProductPrice, err := factory(1000, "big")
  
  if err != nil {
    fmt.Println(err)
    return
  } 
  fmt.Println("Big Product Price:", bigProductPrice.Price())

  smallProductPrice, err := factory(1000, "small")

  if err != nil {
    fmt.Println(err)
    return
  } 
  fmt.Println("Small Product Price:", smallProductPrice.Price())

  mediumProductPrice, err := factory(1000, "medium")

  if err != nil {
    fmt.Println(err)
    return
  } 
  fmt.Println("Medium Product Price:", mediumProductPrice.Price())
}
