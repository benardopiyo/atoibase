package main

import (
	"github.com/01-edu/z01"
)

// AtoiBase converts a numeric string in a given base to an integer value.
func AtoiBase(s string, base string) int {
	// Check if base is valid
	if len(base) < 2 {
		return 0
	}

	// Check if base contains only unique characters and no '+' or '-'
	charMap := make(map[rune]bool)
	for _, c := range base {
		if c == '+' || c == '-' {
			return 0
		}
		if charMap[c] {
			return 0
		}
		charMap[c] = true
	}

	// Create a map from characters to their respective values in base
	valueMap := make(map[rune]int)
	for i, c := range base {
		valueMap[c] = i
	}

	// Convert the string s to an integer in the given base
	result := 0
	for _, c := range s {
		if _, exists := valueMap[c]; !exists {
			return 0 // Invalid character for the base
		}
		result = result*len(base) + valueMap[c]
	}

	return result
}

// Print the integer value of AtoiBase result as characters using z01
func printIntBaseResult(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		var digits []rune
		for n > 0 {
			digits = append(digits, rune('0'+n%10))
			n /= 10
		}
		for i := len(digits) - 1; i >= 0; i-- {
			z01.PrintRune(digits[i])
		}
	}
	z01.PrintRune('\n')
}

func main() {
	// Test cases for AtoiBase
	testCases := []struct {
		s    string
		base string
	}{
		{"125", "0123456789"},
		{"1111101", "01"},
		{"7D", "0123456789ABCDEF"},
		{"uoi", "choumi"},
		{"bbbbbab", "-ab"},
	}

	for _, tc := range testCases {
		result := AtoiBase(tc.s, tc.base)
		printIntBaseResult(result)
	}
}
