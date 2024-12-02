package main

import (
  "fmt"
  "errors"
)

var lowSalaryError error = errors.New("Error")

func imposto(salarioHora float64, horasMes float64) (float64, error) {
  if horasMes < 80 || horasMes < 0  {
    return 0, fmt.Errorf("%w:, the worked cannot have worked less than 80 hours per month")
  }

  salarioTotal := salarioHora * horasMes

  if salarioTotal > 150000 {
    return salarioHora - (salarioHora * 0.10), nil
  } 

  return salarioTotal, nil
}

func main() {
  var salarioHora float64 = 1000
  var horasMes float64 = 100

  salarioTotal, err := imposto(salarioHora, horasMes)
  
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("Salario total recebido é: R$%.2f\n", salarioTotal)
}
