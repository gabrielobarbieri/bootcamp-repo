package main

import "fmt"

func emprestimo(idade int, tempoEmprego int, salario int) {
  if idade > 22 && tempoEmprego > 1 && salario > 100000 {
    fmt.Println("Empréstimo liberado!")
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
}
