package leetcode

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	copy(nums, append(nums[n-k:], nums[:n-k]...))
}
