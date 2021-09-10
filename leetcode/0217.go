package leetcode

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = true
		}
	}
	return false
}
