package learning_tests

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("common input", func(t *testing.T) {
		got := Repeat('A', 5)
		expected := "AAAAA"
		if got != expected {
			t.Errorf("expected :\"%s\", but got: \"%s\"", expected, got)
		}
	})

	t.Run("no repeat", func(t *testing.T) {
		got := Repeat('A', 1)
		expected := "A"
		if got != expected {
			t.Errorf("expected :\"%s\", but got: \"%s\"", expected, got)
		}
	})
}
