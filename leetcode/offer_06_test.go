package leetcode

import (
	"reflect"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *ListNode
	}
	listnode := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: nil}}}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{&listnode}, []int{2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
