package main

import "fmt"

func calcularNotas(notas ...float64) float64 {
  var soma float64

  for _, value := range notas {  
    soma += value  
  }

  return soma / float64(len(notas))
} 

func main() {
  media := calcularNotas(8.5, 9.0, 7.5, 6.0, 8.0)
  
  fmt.Printf("A média do aluno é: %.2f\n", media)
}
