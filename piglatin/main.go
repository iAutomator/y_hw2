package piglatin

import (
	"strings"
)

const ending = "ay"
const vowels = "aeiou"

func countBeginningConsonants(s string) int {
	return strings.IndexAny(strings.ToLower(s), vowels)
}

// Encode encodes a sentence to Pig Latin
func Encode(s string) string {
	words := strings.Split(s, " ")
	for i, w := range words {
		words[i] = encodeWord(w)
	}
	return strings.Join(words, " ")
}

func encodeWord(s string) string {
	nConsonants := countBeginningConsonants(s)
	if nConsonants == 0 {
		return s + ending
	} else if nConsonants == -1 {
		return s
	}

	res := s[nConsonants:] + s[:nConsonants] + ending
	if strings.ToUpper(s[:1]) == s[:1] {
		res = strings.ToUpper(res[:1]) + strings.ToLower(res[1:])
	}
	return res
}
