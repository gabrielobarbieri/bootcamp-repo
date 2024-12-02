package main

import "fmt"

type myError struct {
  errorMsg string
} 

func (e *myError) Error() string {
  return e.errorMsg
}

func imposto(salario float64) (float64, error) {
  if salario < 150000 {
      return 0, &myError{"Error: the salary entered does not reach the taxable minimum"} 
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
  var salario float64 = 1000

  _, err := imposto(salario)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println("Must pay tax")
}
