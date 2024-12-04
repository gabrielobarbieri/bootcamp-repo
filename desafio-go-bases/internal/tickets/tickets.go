package tickets

import (
  "encoding/csv"
  "os"
  "strconv"
  "time"
)

type Ticket struct {
  ID int
  name  string
  email string
  country string
  time string
  randNum int
}

type Tickets struct {
  ticketsList []Ticket } 

func (t *Tickets) Add(ticket Ticket) {
  t.ticketsList = append(t.ticketsList, ticket)
}

func (t *Tickets) LoadFromCSV(filename string) error {
  file, err := os.Open(filename)

  if err != nil {
    return err  
  }

  defer file.Close()

  reader := csv.NewReader(file)
  records, err := reader.ReadAll()

  if err != nil {
    return err
  }

  for _, record := range records {
    id, _ := strconv.Atoi(record[0])
    randNum, _ := strconv.Atoi(record[5])

    ticket := Ticket{
      ID: id,
      name: record[1],
      email: record[2],
      country: record[3],
      time: record[4],
      randNum: randNum,
    }

    t.Add(ticket)
  }
 return nil 
}

// ejemplo 1
func (t *Tickets) GetTotalTickets(destination string) (int, error) {
  count := 0

  for _, ticket := range t.ticketsList { 
    if ticket.country == destination {
      count += 1
    }
  }
  return count, nil
}

// ejemplo 2
func (t *Tickets) GetCountByPeriod(period string) (int, error) {
  // from 7:00 to 12:00
  count := 0
    
  for _, ticket := range t.ticketsList {
    switch period { 
      case "Morning":    
        startTime, _ := time.Parse("15:04", "7:00")         
        endTime, _ :=  time.Parse("15:04", "12:00")
        
        ticketTime, _ := time.Parse("15:04", ticket.time)
        
        if ticketTime.After(startTime) && ticketTime.Before(endTime) {
           count += 1      
        }
      case "Afternoon":
        startTime, _ := time.Parse("15:04", "13:00")         
        endTime, _ :=  time.Parse("15:04", "19:00")

        ticketTime, _ := time.Parse("15:04", ticket.time)

        if ticketTime.After(startTime) && ticketTime.Before(endTime) {
           count += 1      
        }
      case "Night":
        startTime, _ := time.Parse("15:04", "20:00")         
        endTime, _ :=  time.Parse("15:04", "23:00")

        ticketTime, _ := time.Parse("15:04", ticket.time)

        if ticketTime.After(startTime) && ticketTime.Before(endTime) {
           count += 1      
        }
      case "EarlyMorning":
        startTime, _ := time.Parse("15:04", "00:00")         
        endTime, _ :=  time.Parse("15:04", "6:00")

        ticketTime, _ := time.Parse("15:04", ticket.time)

        if ticketTime.After(startTime) && ticketTime.Before(endTime) {
           count += 1      
        }

    }
  }
  return count, nil
}
// ejemplo 3
func (t *Tickets) AverageDestination(destination string) (float64, error) {
  countryCount, _ := t.GetTotalTickets(destination)
  totalTickets := float64(len(t.ticketsList))

  var result float64 = float64(countryCount) / totalTickets

  return result, nil
}
