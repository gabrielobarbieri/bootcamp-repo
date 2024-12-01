package main

import (
	"testing"
)

func TestImposto(t *testing.T) {
  tests := []struct {
    description string
    salario float64
    expected float64
  }{
    {"Teste sem imposto", 20000, 0},
    {"Teste com imposto 17%", 60000, 10200},
    {"Teste com imposto 27%", 170000, 45900},
  }
  
  for _, tt := range tests {
    t.Run(tt.description, func(t *testing.T) {
      result := Imposto(tt.salario) 

      if result != tt.expected {
        t.Errorf("expected: %.2f, got: %.2f", tt.expected, result)
      }
    })
  }
}
