package main

import "fmt"

type Person struct {
  ID int
  Name string
  DateOfBirth string
}

type Employee struct {
  ID int
  Position string
  Person Person
}

func (e Employee) PrintEmployee() {
  fmt.Printf("%s is a %s, born in %s\n", e.Person.Name, e.Position, e.Person.DateOfBirth) 
}

func main() {
  p1 := Person {
    Name: "Gabriel",
    DateOfBirth: "23/01/2004",
  }

  e1 := Employee {
    Position: "Software Engineer",
    Person: p1,
  }

  e1.PrintEmployee()
}
