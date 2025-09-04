package prefixsums

func CalculateProductArrayExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	left_product := left_product(nums)
	right_product := right_product(nums)

	for i := 0; i < n; i++ {
		ans[i] = left_product[i] * right_product[i]
	}

	return ans
}

func left_product(nums []int) []int {
	leftProduct := make([]int, len(nums))
	leftProduct[0] = 1

	for i := 1; i < len(nums); i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}

	return leftProduct
}

func right_product(nums []int) []int {
	rightProduct := make([]int, len(nums))
	rightProduct[len(nums)-1] = 1

	for i := len(nums) - 2; i >= 0; i-- {
		rightProduct[i] = rightProduct[i+1] * nums[i+1]
	}

	return rightProduct
}
