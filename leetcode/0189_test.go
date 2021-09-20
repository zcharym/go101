package leetcode

import (
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums   []int
		k      int
		expect []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"case 1", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}}},
		{"case 2", args{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expect := tt.args.expect
			rotate(tt.args.nums, tt.args.k)
			if !TestSliceEqual(tt.args.nums, expect) {
				t.Errorf("not equal: want:%v, got:%v", expect, tt.args.nums)
			}
		})
	}
}
