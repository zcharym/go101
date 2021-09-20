package leetcode

func ArrayInsert(arr []int, index, value int) []int {
	if len(arr) <= index {
		return append(arr, value)
	}
	arr = append(arr[:index+1], arr[index:]...)
	arr[index] = value
	return arr
}

func TestSliceEqual(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i, v := range nums1 {
		if v != nums2[i] {
			return false
		}
	}
	return true
}
