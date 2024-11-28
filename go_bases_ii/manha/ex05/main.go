package main

import (
  "fmt"
  "errors"
)

func animal(animalName string) (func(int) float64, error) {
  switch animalName {  
    case "Cão":
      return func(quantidade int) float64 {
        return 10 * float64(quantidade)
      }, nil
    case "Gato":
      return func(quantidade int) float64 {
        return 5 * float64(quantidade)
      }, nil
    case "Hamster":
      return func(quantidade int) float64 {
        return 0.25 * float64(quantidade) 
      }, nil
    case "Tarântula":
      return func(quantidade int) float64 {
        return 0.15 * float64(quantidade) 
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
