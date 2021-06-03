package learning_tests

func Add(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
