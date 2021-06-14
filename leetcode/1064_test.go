package leetcode

import (
	"testing"
)

func Test_fixedPoint(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]int{-10, -5, 0, 3, 7}}, 3},
		{"case 2", args{[]int{0, 2, 5, 8, 17}}, 0},
		{"case 3", args{[]int{-10, -5, 3, 4, 7, 9}}, -1},
		{"case 4", args{[]int{-10, -5, -2, 0, 4, 5, 6, 7, 8, 9, 10}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fixedPoint(tt.args.arr); got != tt.want {
				t.Errorf("fixedPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
