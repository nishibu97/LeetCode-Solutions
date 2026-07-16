package main

import (
	"slices"
	"strings"
	"testing"
)

func isPalindrome(s string) bool {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expected     string
		validAnswers []string
	}{
		{
			name:         "例1",
			input:        "babad",
			validAnswers: []string{"bab", "aba"},
		},
		{
			name:     "例2",
			input:    "cbbd",
			expected: "bb",
		},
		{
			name:     "一文字",
			input:    "a",
			expected: "a",
		},
		{
			name:     "全体が回文",
			input:    "racecar",
			expected: "racecar",
		},
		{
			name:     "偶数長の回文",
			input:    "abba",
			expected: "abba",
		},
		{
			name:     "重複文字のみ",
			input:    "aaaa",
			expected: "aaaa",
		},
		{
			name:     "回文なし（最長は一文字）",
			input:    "abc",
			expected: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := longestPalindrome(tt.input)

			if !isPalindrome(actual) {
				t.Errorf("longestPalindrome(%q) = %q; result is not a palindrome", tt.input, actual)
			}

			if !strings.Contains(tt.input, actual) {
				t.Errorf("longestPalindrome(%q) = %q; result is not a substring", tt.input, actual)
			}

			if tt.expected != "" && actual != tt.expected {
				t.Errorf("longestPalindrome(%q) = %q; expected %q", tt.input, actual, tt.expected)
			}

			if len(tt.validAnswers) > 0 && !slices.Contains(tt.validAnswers, actual) {
				t.Errorf(
					"longestPalindrome(%q) = %q; expected one of %v",
					tt.input,
					actual,
					tt.validAnswers,
				)
			}
		})
	}
}
