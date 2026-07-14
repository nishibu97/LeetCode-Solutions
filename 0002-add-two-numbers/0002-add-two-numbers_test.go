package main

import (
	"reflect"
	"testing"
)

// スライスの配列からListNodeを生成するヘルパー関数
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// ListNodeからスライスを復元するヘルパー関数
func listToSlice(node *ListNode) []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func TestAddTwoNumbers(t *testing.T) {
	// テストケースの定義
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "Example 1",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			name:     "Example 2",
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			name:     "Example 3",
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	// 各テストケースを実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 入力値をListNodeに変換
			l1 := createList(tt.l1)
			l2 := createList(tt.l2)

			// テスト対象の関数を実行
			resultNode := addTwoNumbers(l1, l2)

			// 結果をスライスに戻して比較
			result := listToSlice(resultNode)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, result)
			}
		})
	}
}