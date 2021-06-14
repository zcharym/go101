package leetcode

// using binary search
func fixedPoint(arr []int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == mid || arr[mid] > mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if low < len(arr) && arr[low] == low {
		return low
	}
	return -1
}
