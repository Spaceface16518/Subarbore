package translate

func Validate(nums []int) []int {
	elements, lengths := runLengthEncode(nums)
	var invalidIndices []int
	for i := 1; i < len(elements); i++ {
		// If the current element is greater and if the last group of elements
		// had more than one element
		// In other words, only a single element is allowed for subtractive
		// notation; if there is more than one, the string is invalid
		if elements[i] > elements[i-1] && lengths[i-1] > 1 {
			invalidIndices = append(invalidIndices, i)
		}
	}
	return invalidIndices
}

func runLengthEncode(l []int) ([]int, []int) {
	origLen := len(l)
	out := make([]int, 1, origLen)
	lens := make([]int, 1, origLen)

	// Do the first iteration; initalize the first element
	prev := l[0]
	index := 0
	out[0] = prev
	lens[0] = 1
	for i := 1; i < origLen; i++ {
		curr := l[i]
		if curr != prev {
			// Add a new element to the rle
			out = append(out, curr)
			lens = append(lens, 1)
			index++
		} else {
			// Update the current element
			lens[index]++
		}
		prev = curr
	}
	return out, lens
}
