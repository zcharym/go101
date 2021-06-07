package leetcode

import (
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "positive value", args: args{x: 123}, want: 321},
		{name: "positive value", args: args{x: 120}, want: 21},
		{name: "zeo value", args: args{x: 0}, want: 0},
		{name: "negative value", args: args{x: -123}, want: -321},
		{name: "negative value", args: args{x: math.MaxInt32 + 1}, want: 0},
		{name: "negative value", args: args{x: math.MinInt32 - 1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
