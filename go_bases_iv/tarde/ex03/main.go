package main

import (
  "fmt"
  "os"
  "errors"
)

type Client struct {
  name string
  id int
  phoneNumber string
  address string
} 

type Clients struct {
  clientsList []Client
}

func (c *Clients) Add(cl *Client) {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered from panic:", r)
    }
  }()

  if c.Exists(cl.id) {
    panic("Error: client already exists") 
  }
  
  _, err := c.hasZeroValue(cl)
  if err != nil {
    fmt.Println(err)
    return
  }

  c.clientsList = append(c.clientsList, *cl) 
}

func (c *Clients) Exists(id int) bool {
  for _, client := range c.clientsList {
    if client.id == id {
      return true
    }
  }
  return false
} 

func (c *Clients) hasZeroValue(cl *Client) (bool, error) {
  result := cl.name == "" || cl.id == 0 || cl.phoneNumber == "" || cl.address == "" 

  if result {
    return false, errors.New("Error: Client does not have all information!")
  } else {
    return result, nil
  }
}


func main() {
  defer fmt.Println("Several errors were detected at runtime")
  defer fmt.Println("Execution finished")

  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered from panic:", r)  
    } 
  }()

  filename := "customers.txt"
  file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

  if err != nil {
    panic("The indicated file was not found or is damaged")
  } 
  defer file.Close()

  newClient := &Client{"John Doe", 1, "1234567", "Avenue street"}
  anotherClient := &Client{"Jane Roe", 2, "7654321", "Boulevard avenue, 456"}

  cl := &Clients{}

  cl.Add(newClient)
  cl.Add(anotherClient)

  for _, client := range cl.clientsList {
    _, err := fmt.Fprintf(file, "Client: ID = %d, Name = %s, Phone = %s, Address = %s\n",	client.id, client.name, client.phoneNumber, client.address)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}
	fmt.Println("Clients written to file successfully.")
} 

