package sliding_window

func FindSlidingWindow(s1, s2 string) int {

	if len(s2) > len(s1) {
		return 0
	}

	count := 0
	var window_freq [26]int
	var expected_freq [26]int
	left, right := 0, 0

	for _, c1 := range s2 {
		expected_freq[c1-'a']++
	}
	for right < len(s1) {
		window_freq[s1[right]-'a']++
		if right-left+1 == len(s2) {
			// check if two window are equal
			if window_freq == expected_freq {
				count++
			}
			window_freq[s1[left]-'a']--
			left++
		}
		right++
	}
	return count
}
