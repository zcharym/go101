package leetcode

func firstUniqChar(s string) int {
	var alphabet = [26]int{}
	for _, num := range s {
		alphabet[num-'a']++
	}
	for i, ch := range s {
		if alphabet[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
