package tickets_test

import (
  "github.com/bootcamp-go/desafio-go-bases/internal/tickets"
  "fmt"
  "testing"
) 

func createTickets() tickets.Tickets{
  return tickets.Tickets{}
}

func populateTickets(ts *tickets.Tickets) error {
  err :=  ts.LoadFromCSV("../../tickets.csv")
  return err
}

func TestGetTotalTickets(t *testing.T) {
  // tests functions -> for loop -> assert
  ts := createTickets()

  err := populateTickets(&ts)
  if err != nil {  
    t.Fatalf("Could not read from the file: %v", err)
  }


  tests := []struct{
    description string
    country string  
    expected int
  }{
    {"Testing with an existing country", "Finland", 8},
    {"Testing with a non-existing country", "Vatican", 0},
  }
  
  for _, tt := range tests {
    t.Run(tt.description, func(t *testing.T) {
      result, err := ts.GetTotalTickets(tt.country)
      if err != nil {
        fmt.Println("Could not count the ocurrences for that country!")
        return
      }
      
      if result != tt.expected {
        t.Errorf("expected: %d, got: %d", tt.expected, result)
        return 
      }
    })
  }

}

func TestGetCountByPeriod(t *testing.T) {
  ts := createTickets()

  err := populateTickets(&ts)
  if err != nil {  
    t.Fatalf("Could not read from the file: %v", err)
  }

  tests := []struct { 
    description string
    time string
    expected int
  } { 
    {"Testing with an existing time", "Morning", 223},
  }

  for _, tt := range tests {
    t.Run(tt.description, func(t *testing.T) {
      result, _ := ts.GetCountByPeriod(tt.time)

      if result != tt.expected {
        fmt.Printf("expected: %d, got: %d", tt.expected, result)
      }
    })
  }
}

func TestAverageDestination(t *testing.T) {
  ts := createTickets()

  populateTickets(&ts)

  tests := []struct {
    description string
    country string    
    expected float64
  } {
    {"Testando com país", "Finland", 0.008},
  }

  for _, tt := range tests {
    t.Run(tt.description, func(t *testing.T) { 
      result, _ := ts.AverageDestination(tt.country)

      if tt.expected != result {
        fmt.Printf("expected: %.5f, got: %.5f", tt.expected, result)
        return
      }
    })
  }
}
