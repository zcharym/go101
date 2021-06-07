package leetcode

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "positive value with zero", args: args{x: 10}, want: false},
		{name: "negative value", args: args{x: -121}, want: false},
		{name: "negative value", args: args{x: -101}, want: false},
		{name: "positive value", args: args{x: 121}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
