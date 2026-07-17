package main

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		numRows  int
		expected string
	}{
		{
			name:     "例1",
			s:        "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "例2",
			s:        "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			name:     "例3",
			s:        "A",
			numRows:  1,
			expected: "A",
		},
		{
			name:     "numRows >= len(s)",
			s:        "ABC",
			numRows:  1000,
			expected: "ABC",
		},
		{
			name:     "numRows == 1 かつ複数文字",
			s:        "AB",
			numRows:  1,
			expected: "AB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := convert(tt.s, tt.numRows)
			if actual != tt.expected {
				t.Errorf(
					"convert(%q, %d) = %q; expected %q",
					tt.s,
					tt.numRows,
					actual,
					tt.expected,
				)
			}
		})
	}
}
