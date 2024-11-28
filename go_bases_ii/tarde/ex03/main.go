package main

import "fmt"


func calculaSalario(minutos float64, categoria string) float64 {
  var horas float64
  horas = minutos / 60

  switch categoria {
    case "C":
      return 1000 * horas
    case "B":
      var salario float64
      salario = 1500 * horas
      return salario + (salario * 0.2)
    case "A":
      var salario float64
      salario = 3000 * horas
      return salario + (salario * 0.5)
    default:
      return 0.0
  }
}

func main() {
  var minutos float64
  minutos = 120
  categoria := "C"
  
  fmt.Printf("O salário é: US$%.3f\n", calculaSalario(minutos, categoria))

  categoria = "B"

  fmt.Printf("O salário é: US$%.3f\n", calculaSalario(minutos, categoria))

  categoria = "A"

  fmt.Printf("O salário é: US$%.3f\n", calculaSalario(minutos, categoria))
}
