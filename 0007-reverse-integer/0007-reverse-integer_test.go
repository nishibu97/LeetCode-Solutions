package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected int
	}{
		{
			name:     "例1",
			x:        123,
			expected: 321,
		},
		{
			name:     "例2",
			x:        -123,
			expected: -321,
		},
		{
			name:     "例3",
			x:        120,
			expected: 21,
		},
		{
			name:     "正のオーバーフロー",
			x:        1534236469,
			expected: 0,
		},
		{
			name:     "負のオーバーフロー",
			x:        -2147483648,
			expected: 0,
		},
		{
			name:     "0",
			x:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := reverse(tt.x)
			if actual != tt.expected {
				t.Errorf(
					"reverse(%d) = %d; expected %d",
					tt.x,
					actual,
					tt.expected,
				)
			}
		})
	}
}
