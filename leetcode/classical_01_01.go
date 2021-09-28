package leetcode

func isUnique(astr string) bool {
	var alphbet = [26]int{}

	for _, ch := range astr {
		alphbet[ch-'a']++
	}
	for _, num := range alphbet {
		if num > 1 {
			return false
		}
	}
	return true
}
