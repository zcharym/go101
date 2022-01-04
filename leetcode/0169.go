package leetcode

func majorityElement(nums []int) int {
	length := len(nums)
	m := make(map[int]int, length)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v > length/2 {
			return k
		}

	}
	return 0
}

func majorityElementV2(nums []int) int {
	count := 0
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if count < 0 {
			ans = nums[i]
		}

		if nums[i] == ans {
			count++
		} else {
			count--
		}
	}
	return ans
}
