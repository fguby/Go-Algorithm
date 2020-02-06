package leetcode

/**
Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	cur := &ListNode{
		Val:  0,
		Next: nil,
	}

	prev := cur

	carry := 0

	for l1 != nil || l2 != nil {

		x, y := 0, 0

		if l1 != nil {
			x = l1.Val
		}

		if l2 != nil {
			y = l2.Val
		}

		sum := x + y + carry

		carry = sum / 10

		sum = sum % 10

		prev.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}

		prev = prev.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry == 1 {
		prev.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}

	return cur.Next
}

//func adder(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
//
//	for l1.Next != nil || l2.Next != nil {
//		var x, y int
//
//		if l1 != nil {
//			x = l1.Val
//		}
//
//		if l2 != nil {
//			y = l2.Val
//		}
//
//		sum := x + y + carry
//	}
//}
