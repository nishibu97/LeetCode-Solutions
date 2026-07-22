package main

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "例1",
			s:        "42",
			expected: 42,
		},
		{
			name:     "例2: 先頭空白と符号、先行ゼロ",
			s:        "   -042",
			expected: -42,
		},
		{
			name:     "例3: 数字の後に非数字文字",
			s:        "1337c0d3",
			expected: 1337,
		},
		{
			name:     "例4: 数字の後の符号は無視",
			s:        "0-1",
			expected: 0,
		},
		{
			name:     "例5: 数字が全くない",
			s:        "words and 987",
			expected: 0,
		},
		{
			name:     "空文字列",
			s:        "",
			expected: 0,
		},
		{
			name:     "空白のみ",
			s:        "   ",
			expected: 0,
		},
		{
			name:     "正のオーバーフロー",
			s:        "91283472332",
			expected: 2147483647,
		},
		{
			name:     "負のオーバーフロー",
			s:        "-91283472332",
			expected: -2147483648,
		},
		{
			name:     "int32最大値",
			s:        "2147483647",
			expected: 2147483647,
		},
		{
			name:     "int32最小値",
			s:        "-2147483648",
			expected: -2147483648,
		},
		{
			name:     "プラス記号あり",
			s:        "+1",
			expected: 1,
		},
		{
			name:     "符号のみ",
			s:        "+-12",
			expected: 0,
		},
		{
			name:     "小数点は非数字として扱う",
			s:        "3.14159",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := myAtoi(tt.s)
			if actual != tt.expected {
				t.Errorf(
					"myAtoi(%q) = %d; expected %d",
					tt.s,
					actual,
					tt.expected,
				)
			}
		})
	}
}
