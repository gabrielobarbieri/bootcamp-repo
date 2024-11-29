package main

import "fmt"

type Student struct {
  ID int
  name string
  surname string
  date string  
}

var Students []Student

func main() {
  var id int
  var name, surname, date, userContinue string
  

  for {
    fmt.Println("Add a new student? (y) Exit the program (n)")
    fmt.Scanln(&userContinue) 

    if userContinue == "n" {
      break
    }
    
    fmt.Println("Enter student details: ")

    fmt.Print("ID: ")
    fmt.Scanln(&id)

    fmt.Print("Name: ") 
    fmt.Scanln(&name)

    fmt.Print("Surname: ")
    fmt.Scanln(&surname)
    
    fmt.Print("Date: ")
    fmt.Scanln(&date)

    newStudent := Student {
      ID: id,
      name: name,
      surname: surname,
      date: date, 
    }
    
    Students = append(Students, newStudent)
    fmt.Printf("The id of the student is: %d, name: %s, surname: %s, date of birth: %s\n", id, name, surname, date)
    
  }
  
  fmt.Printf("Final list of students: %v", Students)

}
