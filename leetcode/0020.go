package leetcode

var m = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

func isValid(s string) bool {
	var stack []rune
	pushed := false
	for _, ch := range []rune(s) {
		// 左括号压入栈顶
		if ch == 123 || ch == 91 || ch == 40 {
			stack = append(stack, ch)
			pushed = true
		} else if (ch == 125 || ch == 93 || ch == 41) && len(stack) > 0 && m[ch] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			// 单个右括号无匹配
			return false
		}
	}
	if len(stack) == 0 && pushed == true {
		return true
	}
	return false
}
