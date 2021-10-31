package leetcode

import (
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 0},
		{"case 2", args{[]int{1}}, 1},
		{"case 3", args{[]int{0}}, 0},
		{"case 4", args{[]int{-1}}, -1},
		{"case 5", args{[]int{-100000}}, -100000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
