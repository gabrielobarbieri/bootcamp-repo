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

func imposto(salario float64) (float64, error) {
  if salario <= 10000 {
   return 0, lowSalaryError 
  }

	if salario < 50000 {
		return 0, nil
	} else if salario > 50000 && salario < 150000 {
		return salario * 0.17, nil
	} else {
		return salario * 0.27, nil
	}
}

func main() {
  var salary float64  
  salary = 10000

  _, err := imposto(salary) 

  if salary <= 10000 { 
    if errors.Is(err, lowSalaryError) {
      fmt.Println("Error: ", err.Error())
      return
    }
  }
}
