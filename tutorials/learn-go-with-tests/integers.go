package learning_tests

func Add(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func AddAll(arrays ...[]int) []int {
	var sums []int
	for _, nums := range arrays {
		sums = append(sums, Add(nums...))
	}
	return sums
}
