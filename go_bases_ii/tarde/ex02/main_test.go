package main

import (
	"testing"
)

func TestCalcularNotas(t *testing.T) {
  // Create input, create for loop, execute tests

  tests := []struct {
    description string
    notas []float64
    expected float64
  } {
    {"Notas zero", []float64{0, 0, 0}, 0.0},
    {"Notas maiores que zero", []float64{1.0, 4.0, 7.0}, 4},
    {"Notas maiores que zero", []float64{5.0, 5.0, 10.0, 0}, 5.0},
  }

  for _, tt := range tests {
    t.Run(tt.description, func(t *testing.T) {
      got := calcularNotas(tt.notas...)
      if tt.expected != got {
        t.Errorf("expected: %.2f, got: %.2f", tt.expected, got)
      }
    })
  }
}
