package main

import (
  "errors"
  "fmt"
)

func operation(calculo string) (func(...int) int, error) {
  switch calculo { 
    case "minimum":
      return func(notas ...int) int {
        minValue := notas[0]

        for _, nota := range notas {
          if nota < minValue { 
            minValue = nota
          }
        }

       return minValue
      }, nil
    case "average":
      return func(notas ...int) int {
        soma := 0

        for _, nota := range notas {
          soma += nota
        }

        return int(soma/len(notas))
      }, nil
    case "maximum":
      return func(notas ...int) int {
        maxValue := notas[0]   

        for _, nota := range notas {
          if nota > maxValue { 
            maxValue = nota
          }
        }

        return maxValue
      }, nil
    default:
      return nil, errors.New("Calculo não encontrado")
  }
}

func main() {
  const (
     minimum = "minimum"
     average = "average"
     maximum = "maximum"
  )

  minFunc, err := operation(minimum)
  if err != nil {
      fmt.Println(err)
      return
  }

  averageFunc, err := operation(average)
  if err != nil {
      fmt.Println(err)
      return
  }

  maxFunc, err := operation(maximum)
  if err != nil {
      fmt.Println(err)
      return
  }

  minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
  fmt.Printf("Minimum: %d\n", minValue)

  averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
  fmt.Printf("Average: %d\n", averageValue)

  maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
  fmt.Printf("Maximum: %d\n", maxValue)
}
