package main

import "testing"

func TestCalculaSalario(t *testing.T) {
	t.Run("Teste calcular salário", func(t *testing.T) {
		got := calculaSalario(120, "C")
		expected := float64(2000)

		if got != expected {
			t.Errorf("esperado %.2f, obteve %.2f", expected, got)
		}

		got = calculaSalario(180, "B")
		expected = 5400

		if got != expected {
			t.Errorf("esperado %.2f, obteve %.2f", expected, got)
		}

		got = calculaSalario(0, "A")
		expected = 0

		if got != expected {
			t.Errorf("esperado %.2f, obteve %.2f", expected, got)
		}
	})
}
