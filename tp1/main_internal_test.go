package main

import "testing"

func Example_main() {
}

func TestSearch(t *testing.T) {
	prenom := "Elwin"
	want := "Bassaget"
	t.Run(prenom, func(t *testing.T) {
		got := search(prenom)

		if got.nom != want {
			t.Errorf("expected: %q, got: %q", want, got)
		}
	})
}

func TestAdd(t *testing.T) {
	prenom := "Steph"
	nom := "Ramos"
	numero := "06352063"

	want := "Ramos"

	t.Run(prenom, func(t *testing.T) {
		got := addUser(prenom, nom, numero)

		if got.nom != want {
			t.Errorf("expected: %q, got: %q", want, got)
		}
	})
}

func TestModify(t *testing.T) {
	prenom := "Elwin"
	nom := "Steph"
	want := "Steph"
	t.Run(prenom, func(t *testing.T) {
		got := modify(prenom, nom, "", "")

		if got.nom != want {
			t.Errorf("expected: %q, got: %q", want, got)
		}
	})
}
