package leetcode

func findDisappearedNumbers(nums []int) (result []int) {
	var store = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		store[nums[i]-1] = 1
	}
	for index, num := range store {
		if num == 0 {
			result = append(result, index+1)
		}
	}
	return
}
