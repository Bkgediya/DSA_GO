package sliding_window

func LongestSubstring(s string) int {
	max_len := 0
	left, right := 0, 0
	prev_occu := make(map[byte]int)
	for right < len(s) {
		// check if previous occurence of this character is present in the current window
		if prevIndex, exists := prev_occu[s[right]]; exists && prevIndex >= left {
			left = prevIndex + 1
		}

		if right-left+1 > max_len {
			max_len = right - left + 1
		}
		prev_occu[s[right]] = right
		right++
	}

	return max_len
}
