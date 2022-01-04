package leetcode

import "testing"

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candyType []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{candyType: []int{1, 1, 2, 2, 3, 3}}, 3},
		{"case 2", args{candyType: []int{1, 1, 2, 3}}, 2},
		{"case 3", args{candyType: []int{6, 6, 6, 6}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.candyType); got != tt.want {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
