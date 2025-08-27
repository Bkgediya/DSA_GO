package stack

import (
	"unicode"
)

func EvaluateExpression(expression string) int {
	stack := []int{}
	curr_num, sign, res := 0, 1, 0

	for _, ch := range expression {
		if unicode.IsDigit(rune(ch)) {
			curr_num = curr_num*10 + int(ch-'0')
		} else if rune(ch) == '+' || rune(ch) == '-' {
			res += sign * curr_num
			if rune(ch) == '+' {
				sign = 1
			} else if rune(ch) == '-' {
				sign = -1
			}
			curr_num = 0

		} else if rune(ch) == '(' {
			stack = append(stack, res)
			stack = append(stack, sign)
			res = 0
			sign = 1
		} else if rune(ch) == ')' {
			res += sign * curr_num
			curr_num = 0
			res *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return res + curr_num*sign
}
