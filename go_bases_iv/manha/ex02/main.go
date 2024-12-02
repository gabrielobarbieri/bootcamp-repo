package main

import (
  "fmt"
  "errors"
)

type myError struct {
  errorMsg string
} 

func (e *myError) Error() string {
  return e.errorMsg
}

var lowSalaryError *myError = &myError{"Error: salary is less than 10000"}

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
  salary := 10000

  e := lowSalaryError 

  if salary <= 10000 { 
    fmt.Println(errors.Is(e, lowSalaryError))
    if errors.Is(e, lowSalaryError) {
      fmt.Println("Error: ", e.Error())
      return
    }
  }
}
