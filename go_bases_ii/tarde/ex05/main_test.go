package main

import "testing"

func TestAnimal(t *testing.T) {
  // Create tests table -> for loop -> assert
  tests := []struct {
    description string
    animalName string
    quantidade float64
    expected float64
  }{
    {"Dog calculation", "Cão", 10.0, 100},
    {"Cat calculation", "Gato", 0.5, 2.5},
    {"Hamster calculation", "Hamster", 20.0, 5},
    {"Tarântula calculation", "Tarântula", 30.0, 4.5},
  }
  
  for _, tt := range tests {
    t.Run(tt.description, func(t *testing.T) {
      animalCalc, _ := animal(tt.animalName)
      
      got := animalCalc(tt.quantidade)
      if tt.expected != got {
        t.Errorf("expected: %.2f, got: %.2f", tt.expected, got)
        return
      }
    })
  }
} 
