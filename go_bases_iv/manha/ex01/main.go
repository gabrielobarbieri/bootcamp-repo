package main

import "fmt"

type myError struct {
  errorMsg string
} 

func (e *myError) Error() string {
  return e.errorMsg
}

func imposto(salario float64) float64 {
	if salario < 50000 {
		return 0
	} else if salario > 50000 && salario < 150000 {
		return salario * 0.17
	} else {
		return salario * 0.27
	}
}

func main() {
  salary := 1000

  e := myError{"Error: the salary entered does not reach the taxable minimum"}

  if salary < 150000 {
    fmt.Println(e.Error())
    return
  } 

  fmt.Println("Must pay tax")
}
