package main

import "fmt"

type Product interface {
  price() float64
}

type smallProduct struct {
  productPrice float64
}

type mediumProduct struct {
  productPrice float64
}

type bigProduct struct {
  productPrice float64
}

// criar func factory -> métodos -> criar interface product (com metodo price)
func (s smallProduct) price() float64 {
  return s.productPrice 
}

func (m mediumProduct) price() float64 {
  return m.productPrice + ((m.productPrice * 0.03) * 2)
}

func (b bigProduct) price() float64 {
  return b.productPrice + (b.productPrice * 0.06) + 2500
}

func factory(productType string, productPrice float64) float64{
  switch productType {
    case "small":
      productInstance := smallProduct{productPrice}
      return productInstance.price() 
    case "medium":
      productInstance := mediumProduct{productPrice}
      return productInstance.price()
    case "big":
      productInstance := bigProduct{productPrice}
      return productInstance.price()
    default: 
      return 0.0 
  }
}

func main() {
  bigProductPrice := factory("big", 1000)
  
  fmt.Println(bigProductPrice)

  smallProductPrice := factory("small", 1000)

  fmt.Println(smallProductPrice)

  mediumProductPrice := factory("medium", 1000)

  fmt.Println(mediumProductPrice)
}
