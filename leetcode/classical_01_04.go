package leetcode

// 回文串的判断依据：字母集合中字母成对存在或者仅存在一个数量为一的字母
func canPermutePalindrome(s string) bool {
	var m = make(map[rune]int, 26)

	for _, v := range s {
		m[v] += 1
	}

	odd := 0
	for _, v := range m {
		if v%2 != 0 {
			odd++
		}
	}
	return odd == 0 || odd == 1

}
