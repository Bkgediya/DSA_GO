package stack

func IsValidParenthesis(s string) bool {
	parenthesis_map := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []rune{}

	for _, val := range s {
		if _, exists := parenthesis_map[rune(val)]; exists {
			stack = append(stack, rune(val))
		} else {
			if len(stack) > 0 && parenthesis_map[stack[len(stack)-1]] == rune(val) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
