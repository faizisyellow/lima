package utils

import "testing"

func TestUtils(t *testing.T) {

	t.Run("should upper first character", func(t *testing.T) {

		input := "lizzy"
		expected := "Lizzy"

		res := ToUpperFirst(input)
		if res != expected {
			t.Errorf("expected (%v) but got %v", expected, res)
		}
	})

	t.Run("should return empty string", func(t *testing.T) {

		input := ""
		expected := ""

		res := ToUpperFirst(input)
		if res != expected {
			t.Errorf("expected (%v) but got %v", expected, res)
		}
	})
}
