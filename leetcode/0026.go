package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	} else {
		i := 1
		j := 0
		for ; i < len(nums); i++ {
			if nums[i] != nums[j] {
				j += 1
				nums[j] = nums[i]
			}
		}
		nums = nums[:j+1]
		return len(nums)
	}
}
