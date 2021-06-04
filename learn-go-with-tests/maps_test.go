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
		expected := "can't find the word"
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertStrings(t, err.Error(), expected)
	})

	t.Run("add already added word", func(t *testing.T) {
		err := dict.Add("test", "definition for test")
		expected := "word already added"
		if err == nil {
			t.Fatal("expected to get an erro")
		}
		assertStrings(t, err.Error(), expected)
	})

	t.Run("add  word", func(t *testing.T) {
		dict.Add("test add", "definition for test")
		expected := "definition for test"
		got, _ := dict.Search("test add")
		assertStrings(t, got, expected)
	})

	t.Run("test  update", func(t *testing.T) {
		newText := "newText item"
		dict.Update("test", newText)
		expected, _ := dict.Search("test")
		assertStrings(t, expected, newText)
	})

	t.Run("test  delete", func(t *testing.T) {
		dict.Delete("test")
		_, err := dict.Search("test")
		if err != ErrorNotFound {
			t.Errorf("Expected '%s' to be deleted", "test")
		}
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
