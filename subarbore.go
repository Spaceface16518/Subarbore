package subarbore

import (
	"subarbore/sum"
	"subarbore/translate"
)

func ParseRegular(numeric string) int {
	return sum.Sum(translate.String(numeric))
}

func ParseRegularChecked(numeric string) (int, []int) {
	translated := translate.String(numeric)
	invalid := translate.Validate(translated)
	sum := sum.Sum(translated)
	return sum, invalid
}

func ParseSubtractive(numeric string) int {
	return sum.SubtractiveSum(translate.String(numeric))
}

func ParseSubtractiveChecked(numeric string) (int, []int) {
	translated := translate.String(numeric)
	invalid := translate.Validate(translated)
	sum := sum.SubtractiveSum(translated)
	return sum, invalid
}

var globalSubtractive = true

func Parse(numeric string) int {
	if globalSubtractive {
		return ParseSubtractive(numeric)
	} else {
		return ParseRegular(numeric)
	}
}

func ParseChecked(numeric string) (int, []int) {
	if globalSubtractive {
		return ParseSubtractiveChecked(numeric)
	} else {
		return ParseRegularChecked(numeric)
	}
}

func SetGlobalDefaultSubtractive() {
	globalSubtractive = true
}

func SetGlobalDefaultRegular() {
	globalSubtractive = false
}

func GlobalDefaultSubtractive() bool {
	return globalSubtractive
}

func GlobalDefaultRegular() bool {
	return !globalSubtractive
}
