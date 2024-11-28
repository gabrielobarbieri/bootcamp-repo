package main

import (
	"testing"
)

func TestCalcularNotas(t *testing.T) {
	t.Run("Teste calcular média", func(t *testing.T) {
		got := calcularNotas(8.5, 9.0, 7.5, 6.0, 8.0)
		expected := 7.80

		if got != expected {
			t.Errorf("esperado %.2f, obteve %.2f", expected, got)
		}

		got = calcularNotas(2.0, 3.5, 9.0, 5.0)
		expected = 4.875

		if got != expected {
			t.Errorf("esperado %.3f, obteve %.3f", expected, got)
		}
	})
}
