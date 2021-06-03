package learning_tests

import (
	"reflect"
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

func TestAddAll(t *testing.T) {
	t.Run("two arrays", func(t *testing.T) {
		got := AddAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		// NOTE
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected '%d' but got '%d", expected, got)
		}
	})
}
