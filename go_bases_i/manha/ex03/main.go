package main

import "fmt"

func main() {
  var firstName string = "John"
  var lastName string = "Doe"
  var age int = 25
  lastName = "Jon" 
  var driverLicense bool = false
  var personHeight float32 = 1.70
  childNumber := 2

  fmt.Printf("First name: %s\n", firstName)
  fmt.Printf("Last name: %s\n", lastName)
  fmt.Printf("Age: %d\n", age)
  fmt.Printf("Driver license: %t\n", driverLicense)
  fmt.Printf("First name: %f\n", personHeight)
  fmt.Printf("Children: %d\n", childNumber)
}
