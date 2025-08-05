package arrays

import "unicode"

// time complexity O(n)
// space complexity O(1)
func Palindrome(value string) bool {
	i, j := 0, len(value)-1

	for i < j {

		for i < j && is_non_alphanumerical(value[i]) {
			i++
		}
		for i < j && is_non_alphanumerical(value[j]) {
			j--
		}

		if unicode.ToLower(rune(value[i])) != unicode.ToLower(rune(value[j])) {
			return false

		}
		i++
		j--
	}
	return true
}

func is_non_alphanumerical(char byte) bool {
	return !unicode.IsDigit(rune(char)) && !unicode.IsLetter(rune(char))
}
