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

// better solution:
func findDisappearedNumbersV2(nums []int) []int {
	n := len(nums)
	for _, num := range nums {
		nums[(num-1)%n] += n
	}
	res := make([]int, 0, n)
	for index, num := range nums {
		if num <= n {
			res = append(res, index+1)
		}
	}
	return res
}
