package main

import "fmt"

func emprestimo(idade int, tempoEmprego int, salario int) {
  if idade > 22 && tempoEmprego > 1 {
    if salario > 100000 {
      fmt.Println("Empréstimo liberado sem juros!")
    } else {
      fmt.Println("Empréstimo liberado com juros!")   
    }
  } else { 
    fmt.Println("Empréstimo recusado!")
  }
}

func main() {
  idade := 23
  tempoEmprego := 3
  salario := 200000

  emprestimo(idade, tempoEmprego, salario)

  idade = 20

  emprestimo(idade, tempoEmprego, salario)

  idade = 23
  salario = 90000

  emprestimo(idade, tempoEmprego, salario)
}
