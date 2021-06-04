package learning_tests

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		expected := "this is a test"

		assertStrings(t, got, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		want := "can't find the word"
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t *testing.T, got, expected string) {
	//  Helper marks the calling function as a test helper function
	//  When printing file and line information,
	//  that function will be skipped
	t.Helper()

	if got != expected {
		t.Errorf("got '%s' want '%s'", got, expected)
	}
}
