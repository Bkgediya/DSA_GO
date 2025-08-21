package sliding_window

func FindLongestUniformSubstring(s string, k int) int {
	max_len, highest_freq_char := 0, 0
	left, right := 0, 0
	freq := make(map[byte]int)

	for right < len(s) {
		freq[s[right]] = freq[s[right]] + 1
		highest_freq_char = max(highest_freq_char, freq[s[right]])
		number_of_character_to_replace := (right - left + 1) - highest_freq_char
		if number_of_character_to_replace > k {
			freq[s[right]] = freq[s[right]] - 1
			left++
		}
		max_len = right - left + 1
		right++
	}
	return max_len
}
