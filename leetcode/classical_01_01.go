package leetcode

func isUnique(astr string) bool {
	if len(astr) < 2 {
		return true
	}

	m := make(map[rune]bool)
	for _, v := range astr {
		if _, ok := m[v]; ok {
			return false
		} else {
			m[v] = true
		}
	}
	return true
}
