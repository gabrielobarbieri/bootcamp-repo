package main

import (
	"testing"
)

func TestImposto(t *testing.T) {
	got := imposto(20000)
	expected := float64(0)

	t.Run("Teste nenhum imposto", func(t *testing.T) {
		if got != expected {
			t.Errorf("esperado %.2f, obteve %.2f", expected, got)
		}
	})

	got = imposto(55000)
	expected = 9350

	t.Run("Teste 17% imposto", func(t *testing.T) {
		if got != expected {
			t.Errorf("esperado %.2f, obteve %.2f", expected, got)
		}
	})

	got = imposto(200000)
	expected = 54000

	t.Run("Teste 27% imposto", func(t *testing.T) {
		if got != expected {
			t.Errorf("esperado %.2f, obteve %.2f", expected, got)
		}
	})
}
