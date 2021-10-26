package leetcode

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	if n == 0 || n == 1 {
		return n
	}
	ans := 0

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	for low, fast := 0, -1; low < n; low++ {
		if low != 0 {
			delete(m, s[low-1])
		}
		for fast+1 < n && m[s[fast+1]] == 0 {
			m[s[fast+1]]++
			fast++
		}
		ans = max(ans, fast-low+1)
	}
	return ans
}
