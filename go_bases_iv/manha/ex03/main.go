package main

import (
  "fmt"
  "errors"
)

var lowSalaryError error = errors.New("Error: salary is less than 10000")

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
  var salary float64 = 10000

  _, err := imposto(salary)

  if salary <= 10000 { 
    fmt.Println(errors.Is(err, lowSalaryError))

    if errors.Is(err, lowSalaryError) {
      fmt.Println("Error: ", err)
      return
    }
  }
}
