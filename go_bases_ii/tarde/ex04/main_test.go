package main

import (
	"testing"
)

func TestOperation(t *testing.T) {
  // Write tests table -> for loop -> assert
  tests := []struct {
    description string
    calculation string
    inputs []int
    expected int
  }{
    {"Minimum calculation", "minimum", []int{10, 7, 9, 5, 2}, 2},
    {"Average calculation", "average", []int{2, 3, 3, 4, 1, 2, 4, 5}, 3},
    {"Maximum calculation", "maximum", []int{10, 7, 9, 5, 2}, 10},
  } 

  for _, tt := range tests {
    t.Run(tt.description, func(t *testing.T) {
      calcFunc, _ := operation(tt.calculation)
      got := calcFunc(tt.inputs...)

      if got != tt.expected {
        t.Errorf("expected: %d, got: %d", tt.expected, got)
        return
      }
    })
  } 
}
