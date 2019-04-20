package sum

// Sum is a simple helper function that sums a slice of integers
func Sum(nums []int) (sum int) {
	for _, i := range nums {
		sum += i
	}
	return
}
