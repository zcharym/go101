package leetcode

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{s: "abcabcbb"}, 3},
		{"case 2", args{s: "bbbbbbbb"}, 1},
		{"case 3", args{s: "pwwkew"}, 3},
		{"case 4", args{s: ""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
