package main

import (
	"testing"
)

func Test_fibonacci(t *testing.T) {
	tests := []struct {
		name string
		want func() int
	}{
		{name: "seq 0", want: func() int {
			return 0
		}}, {name: "seq 1", want: func() int {
			return 1
		}}, {name: "seq 2", want: func() int {
			return 1
		}}, {name: "seq 3", want: func() int {
			return 2
		}}, {name: "seq 4", want: func() int {
			return 3
		}}, {name: "seq 5", want: func() int {
			return 5
		}}, {name: "seq 6", want: func() int {
			return 8
		}},
	}
	var f = fibonacci()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(); got != tt.want() {
				t.Errorf("fibonacci() = %v, want %d", got, tt.want())
			}
		})
	}
}
