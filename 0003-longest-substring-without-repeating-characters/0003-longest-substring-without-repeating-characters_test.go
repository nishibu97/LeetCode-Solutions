package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "例1", input: "abcabcbb", expected: 3},
		{name: "例2", input: "bbbbb", expected: 1},
		{name: "例3", input: "pwwkew", expected: 3},
		{name: "空文字列", input: "", expected: 0},
		{name: "一文字", input: "a", expected: 1},
		{name: "重複なし", input: "abcdef", expected: 6},
		{name: "左端より前の重複", input: "abba", expected: 2},
		{name: "空白を含む", input: "a b c a", expected: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := lengthOfLongestSubstring(tt.input)
			if actual != tt.expected {
				t.Errorf(
					"lengthOfLongestSubstring(%q) = %d; expected %d",
					tt.input,
					actual,
					tt.expected,
				)
			}
		})
	}
}
