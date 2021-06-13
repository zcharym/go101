package leetcode

func replaceSpace(s string) string {
	newStr := make([]rune, len(s)*3)
	i := 0
	for _, ch := range s {
		if ch != ' ' {
			newStr[i] = ch
			i++
		} else {
			newStr[i] = '%'
			newStr[i+1] = '2'
			newStr[i+2] = '0'
			i += 3
		}
	}
	return string(newStr[:i])
}
