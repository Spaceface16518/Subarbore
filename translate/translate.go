package translate

import "strings"

// Rune takes a single character that it assumes to be valid and looks up it's equivalent in the character table. If the character provided is invalid, Rune will return 0. Any non-zero value is assumed to be valid.
func Rune(char rune) int {
	return table[char]
}

// String takes a full string, converts it to all uppercase characters and translates every character. The returned slice skips any invalid characters, but will have the same capacity as the length of the string.
func String(s string) []int {
	s = strings.ToUpper(s)
	b := make([]int, 0, len(s))
	for _, c := range s {
		translated := Rune(c)
		if translated != 0 {
			b = append(b, translated)
		}
	}
	return b
}
