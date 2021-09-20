package leetcode

import (
	"sort"
)

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = v * v
	}
	sort.Ints(ans)
	return ans
}
