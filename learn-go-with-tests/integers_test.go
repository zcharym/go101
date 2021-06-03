package learning_tests

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("empty input", func(t *testing.T) {

		sum := Add()
		expected := 0
		if sum != expected {
			t.Errorf("expected '%d' but got '%d", expected, sum)
		}
	})

	t.Run("two integers", func(t *testing.T) {
		sum := Add(2, 100)
		expected := 102
		if sum != expected {
			t.Errorf("expected '%d' but got '%d", expected, sum)
		}
	})

	t.Run("multiple integers", func(t *testing.T) {
		sum := Add(2, 100, -10, 100, 20)
		expected := 212
		if sum != expected {
			t.Errorf("expected '%d' but got '%d", expected, sum)
		}
	})
}
