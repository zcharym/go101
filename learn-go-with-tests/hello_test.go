package main

import "testing"

func TestHello(t *testing.T) {
	got := greeting("Alex")
	want := "hello,Alex"
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
