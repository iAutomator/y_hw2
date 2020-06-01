package customencode

import (
	"strconv"
	"strings"
)

var vowels = "aeiou"

// Encode encodes a string by replacing vowels
func Encode(s string) string {
	s = strings.ToLower(s)
	return strings.Map(func(r rune) rune {
		i := strings.IndexRune(vowels, r)
		if i < 0 {
			return r
		}
		c := strconv.Itoa(i + 1)
		return []rune(c)[0]
	}, s)
}

// Decode decodes a string by returning replaced vowels
func Decode(s string) string {
	s = strings.ToLower(s)
	return strings.Map(func(r rune) rune {
		i, err := strconv.Atoi(string(r))
		if err != nil || i > len(vowels) {
			return r
		}
		return rune(vowels[i-1])
	}, s)
}
