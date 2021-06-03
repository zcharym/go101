package learning_tests

import "testing"

func TestHello(t *testing.T) {
	t.Run("greeting to Alex", func(t *testing.T) {
		got := greeting("Alex", "Spanish")
		want := "hello, Alex in Spanish"
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("empty input for greeting", func(t *testing.T) {
		got := greeting("", "")
		want := "hello, Tom in Chinese"
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})
}
