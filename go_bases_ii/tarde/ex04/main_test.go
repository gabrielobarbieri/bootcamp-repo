package main

import (
	"testing"
)

func TestMinFunc(t *testing.T) {
	minFunc, err := operation("minimum")

	if err != nil {
		t.Fatal(err)
	}
	t.Run("Teste minimum", func(t *testing.T) {
		got := minFunc(1, 2, 3, 4, 5)
		expected := 1

		if got != expected {
			t.Errorf("esperado %d, obteve %d", expected, got)
		}

		got = minFunc(300, 25, 10, 5, 100)
		expected = 5

		if got != expected {
			t.Errorf("esperado %d, obteve %d", expected, got)
		}
	})

	averageFunc, err := operation("average")

	if err != nil {
		t.Fatal(err)
	}
	t.Run("Teste average", func(t *testing.T) {
		got := averageFunc(1, 2, 3, 4, 5)
		expected := 3

		if got != expected {
			t.Errorf("esperado %d, obteve %d", expected, got)
		}

		got = averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		expected = 3

		if got != expected {
			t.Errorf("esperado %d, obteve %d", expected, got)
		}
	})

	maxFunc, err := operation("maximum")

	if err != nil {
		t.Fatal(err)
	}
	t.Run("Teste maximum", func(t *testing.T) {
		got := maxFunc(1, 2, 3, 4, 5)
		expected := 5

		if got != expected {
			t.Errorf("esperado %d, obteve %d", expected, got)
		}

		got = maxFunc(2, 3, 90, 24, 1, 15, 2, 30)
		expected = 90

		if got != expected {
			t.Errorf("esperado %d, obteve %d", expected, got)
		}
	})
}
