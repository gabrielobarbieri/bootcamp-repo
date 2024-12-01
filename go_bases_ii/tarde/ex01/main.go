package main

import "fmt"

func Imposto(salario float64) float64 {
	if salario < 50000 {
		return 0
	} else if salario > 50000 && salario < 150000 {
		return salario * 0.17
	} else {
		return salario * 0.27
	}
}

func main() {
	var salario float64

	salario = 20000

	fmt.Println(Imposto(salario))

	salario = 55000

	fmt.Println(Imposto(salario))

	salario = 200000

	fmt.Println(Imposto(salario))
}
