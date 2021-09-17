package leetcode

func ArrayInsert(arr []int, index, value int) []int {
	if len(arr) <= index {
		return append(arr, value)
	}
	arr = append(arr[:index+1], arr[index:]...)
	arr[index] = value
	return arr
}
