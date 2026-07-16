package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "例1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "例2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "例3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf(
					"twoSum(%v, %d) = %v; expected %v",
					tt.nums,
					tt.target,
					actual,
					tt.expected,
				)
			}
		})
	}
}
