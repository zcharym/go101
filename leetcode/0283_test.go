package leetcode

import (
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		expect []int
	}{
		{"case 1",args{[]int{0,1,0,3,12}},[]int{1,3,12,0,0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int,len(tt.args.nums))
			moveZeroes(tt.args.nums)
			copy(arr,tt.args.nums )
			if !TestSliceEqual(arr,tt.expect){
				t.Errorf("Not Equal,want: %v got: %v",tt.expect,arr)
			}
		})
	}
}
