package leetcode

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1",
			args{&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: nil}}}, 2},
			&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
