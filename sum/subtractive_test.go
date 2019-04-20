package sum

import "testing"

func TestSubtractiveSumNil(t *testing.T) {
	var s []int
	expected := 0
	actual := SubtractiveSum(s)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestSubtractiveSumRegular(t *testing.T) {
	s := []int{10, 10, 5, 1}
	expected := 26
	actual := SubtractiveSum(s)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestSubtractiveSumSubtractive(t *testing.T) {
	s := []int{10, 10, 1, 5}
	expected := 24
	actual := SubtractiveSum(s)
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestSubtractiveSumInvalid(t *testing.T) {
	s := []int{10, 10, 1, 1, 5}
	expected := 25
	actual := SubtractiveSum(s)
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func BenchmarkSubtractiveSum0(b *testing.B) {
	var s []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SubtractiveSum(s)
	}
}

func BenchmarkSubtractiveSum1(b *testing.B) {
	s := []int{5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SubtractiveSum(s)
	}
}

func BenchmarkSubtractiveSum2(b *testing.B) {
	s := []int{10, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SubtractiveSum(s)
	}
}

func BenchmarkSubtractiveSum3(b *testing.B) {
	s := []int{10, 1, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SubtractiveSum(s)
	}
}

func BenchmarkSubtractiveSum5(b *testing.B) {
	s := []int{50, 10, 10, 1, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SubtractiveSum(s)
	}
}

func BenchmarkSubtractiveSum10(b *testing.B) {
	s := []int{1000, 1000, 1000, 1000, 1000, 1000, 100, 500, 5, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SubtractiveSum(s)
	}
}
