package main

import "testing"

func TestAnimal(t *testing.T) {
	animalDog, msg := animal("Cão")
	if msg != nil {
		t.Fatal(msg)
	}

	animalCat, msg := animal("Gato")
	if msg != nil {
		t.Fatal(msg)
	}

	animalHamster, msg := animal("Hamster")
	if msg != nil {
		t.Fatal(msg)
	}

	animalTarantula, msg := animal("Tarântula")
	if msg != nil {
		t.Fatal(msg)
	}

	got := animalDog(10)
	expected := float64(100)

	if got != expected {
		t.Errorf("esperado %.2f, obteve %.2f", expected, got)
	}

	got = animalCat(10)
	expected = float64(50)

	if got != expected {
		t.Errorf("esperado %.2f, obteve %.2f", expected, got)
	}

	got = animalHamster(10)
	expected = 2.5

	if got != expected {
		t.Errorf("esperado %.2f, obteve %.2f", expected, got)
	}

	got = animalTarantula(10)
	expected = 1.5

	if got != expected {
		t.Errorf("esperado %.2f, obteve %.2f", expected, got)
	}
}
