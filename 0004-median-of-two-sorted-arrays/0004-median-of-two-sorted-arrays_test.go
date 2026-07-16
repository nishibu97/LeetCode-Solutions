package main

import (
	"math"
	"testing"
)

const epsilon = 1e-5

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			name:     "例1",
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2.0,
		},
		{
			name:     "例2",
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
		{
			name:     "nums1が空",
			nums1:    []int{},
			nums2:    []int{1},
			expected: 1.0,
		},
		{
			name:     "nums2が空",
			nums1:    []int{2},
			nums2:    []int{},
			expected: 2.0,
		},
		{
			name:     "同じ値のみ",
			nums1:    []int{1, 1},
			nums2:    []int{1, 1},
			expected: 1.0,
		},
		{
			name:     "長さが異なる配列",
			nums1:    []int{1, 2},
			nums2:    []int{3, 4, 5, 6},
			expected: 3.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findMedianSortedArrays(tt.nums1, tt.nums2)
			if math.Abs(actual-tt.expected) > epsilon {
				t.Errorf(
					"findMedianSortedArrays(%v, %v) = %v; expected %v",
					tt.nums1,
					tt.nums2,
					actual,
					tt.expected,
				)
			}
		})
	}
}
