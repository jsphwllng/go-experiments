package main

import (
	"strings"
)

// OnlyVowels returns the vowels from a string
func OnlyVowels(s string) string {
	var answer string
	for _, char := range s {
		if strings.Contains("aeiouAEIOU", string(char)) {
			answer += string(char)
		}
	}
	return strings.ToLower(answer)
}
