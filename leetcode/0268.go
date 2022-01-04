package leetcode

func missingNumber(nums []int) int {
	length := len(nums)
	total := length * (length + 1) / 2
	for _, v := range nums {
		total -= v
	}
	return total
}
