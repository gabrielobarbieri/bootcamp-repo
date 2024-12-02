package main

import (
  "fmt"
  "errors"
)

func animal(animalName string) (func(float64) float64, error) {
  switch animalName {  
    case "Cão":
      return func(quantidade float64) float64 {
        return 10 * quantidade
      }, nil
    case "Gato":
      return func(quantidade float64) float64 {
        return 5 * quantidade
      }, nil
    case "Hamster":
      return func(quantidade float64) float64 {
        return 0.25 * quantidade
      }, nil
    case "Tarântula":
      return func(quantidade float64) float64 {
        return 0.15 * quantidade
      }, nil
    default:
      return nil, errors.New("Animal não encontrado!")
  }
}

func main() {
  const (
     dog = "Cão"
     cat = "Gato"
  )
  
  animalDog, msg := animal(dog)
  if msg != nil {
    fmt.Println(msg)
  }

  animalCat, msg := animal(cat)
  if msg != nil {
    fmt.Println(msg)
  }

  var amount float64

  amount += animalDog(10)
  fmt.Printf("Quantidade de alimento necessário em KG: %.2f\n", amount)
  amount += animalCat(10)
  fmt.Printf("Quantidade de alimento necessário em KG: %.2f\n", amount)
}
