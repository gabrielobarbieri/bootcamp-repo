package main

import "fmt"

func main() {
  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
  
  fmt.Println(employees["Benjamin"])

  count := 0

  for _, age := range employees {
    if age > 21 { 
      count += 1 
    }
  }
  
  fmt.Println(count)

  employees["newEmployee"] = 30

  delete(employees, "Pedro")

  fmt.Println(employees)
}
