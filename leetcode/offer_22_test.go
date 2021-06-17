package leetcode

import (
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case 1", args{&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4,
			Next: &ListNode{Val: 5, Next: nil}}}}, 2}, &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEnd(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
