package linkedlist

func IsHappyNumber(n int) bool {
	slow, fast := n, n
	for {
		slow = findNextNumber(slow)
		fast = findNextNumber(findNextNumber(fast))
		if slow == fast {
			return false
		} else if fast == 1 {
			return true
		}
	}
}

func findNextNumber(n int) int {
	next_num := 0
	for n > 0 {
		digit := n % 10
		n /= 10
		next_num += digit * digit
	}
	return next_num
}
