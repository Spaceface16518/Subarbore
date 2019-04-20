package sum

// SubtractiveSum is a helper function that sums a slice of integers accounting for subtractive notation.
func SubtractiveSum(nums []int) (sum int) {
	max := len(nums) - 1
	if max < 0 {
		return 0
	} else if max == 0 {
		return nums[0]
	}
	for i := 0; i < max; i++ {
		curr := nums[i]
		if curr < nums[i+1] {
			sum -= curr
		} else {
			sum += curr
		}
	}
	sum += nums[max]
	return
}
