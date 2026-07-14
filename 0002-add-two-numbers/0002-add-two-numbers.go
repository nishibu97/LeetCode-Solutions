package main

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // 1. 空のノードを作り、そのメモリアドレスお(&)をポインタとしてdummyに代入
    dummy := &ListNode{}

    // 現在の位置を示すポインタ
    current := dummy

    // 繰り上がりを保持する変数
    carry := 0

    for l1 != nil || l2 != nil || carry > 0 {
        v1, v2 := 0, 0
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }

        sum := v1 + v2 + carry
        carry = sum /10

        current.Next = &ListNode{Val: sum % 10}

        current = current.Next
    }

    return dummy.Next
}