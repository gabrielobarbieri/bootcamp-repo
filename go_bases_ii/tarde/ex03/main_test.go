package main

import "testing"

func TestCalculaSalario(t *testing.T) {
  // Create tests table ->  for loop ->  assert behavior
  tests := []struct { 
    description string
    minutos float64
    categoria string
    expected float64
  }{
    {"Teste categoria C", 120, "C", 2000},
    {"Teste categoria B", 30, "B", 900},
    {"Teste categoria A", 60, "A", 4500},
  }

  for _, tt := range tests {
    t.Run(tt.description, func(t *testing.T) {
      got := calculaSalario(tt.minutos, tt.categoria)

      if tt.expected != got {
        t.Errorf("expected: %.2f, got: %.2f", tt.expected, got)
        return 
      }
    })
  }
}
