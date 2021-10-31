package learning_tests

func Repeat(t byte, times int) string {
	bytes := make([]byte, times)
	for i := 0; i < times; i++ {
		bytes[i] = t
	}
	return string(bytes[:])
}
