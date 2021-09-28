package leetcode

func CheckPermutation(s1 string, s2 string) bool {
	var dict1, dict2 = [26]int{}, [26]int{}

	if len(s1) != len(s2) {
		return false
	}

	for _, ch := range s1 {
		dict1[ch-'a']++
	}

	for _, ch := range s2 {
		dict2[ch-'a']++
	}
	for i := 0; i < 26; i++ {
		if dict1[i] != dict2[i] {
			return false
		}
	}

	return true
}

func CheckPermutationV2(s1 string, s2 string) bool {
	var dict = [26]int{}

	if len(s1) != len(s2) {
		return false
	}

	for _, v := range s1 {
		dict[v-'a']++
	}

	for _, v := range s2 {
		dict[v-'a']--
	}

	for _, v := range dict {
		if v > 0 {
			return false
		}
	}

	return true
}
