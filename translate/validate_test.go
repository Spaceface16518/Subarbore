package translate

import "testing"

func TestRunLengthEncode(t *testing.T) {
	xs := []int{1, 1, 2, 2, 3, 3, 3, 5, 5, 5, 5, 2, 3}
	expectedOut, expectedLens := []int{1, 2, 3, 5, 2, 3}, []int{2, 2, 3, 4, 1, 1}
	actualOut, actualLens := runLengthEncode(xs)

	if len(expectedOut) != len(actualOut) || len(expectedLens) != len(actualLens) {
		t.Error("Lengths did not match")
	}
	for i := 0; i < len(expectedOut); i++ {
		if actualOut[i] != expectedOut[i] || actualLens[i] != expectedLens[i] {
			t.Errorf("Mismatch encountered: expected out: %v, actual out: %v; expected length: %v, actual length: %v; all at index %d", expectedOut[i], actualOut[i], expectedLens[i], actualLens[i], i)
		}
	}
}
