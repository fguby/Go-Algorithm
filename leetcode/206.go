package leetcode

/**
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
**/

// type ListNode struct {
// 	Value int
// 	Next  *ListNode
// }

//递归解法
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	node := reverseList(head.Next)

	head.Next.Next = head

	head.Next = nil

	return node
}

//遍历解法
func reverseList1(head *ListNode) *ListNode {

	var prev *ListNode
	curr := head

	for curr != nil {
		tempNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = tempNode
	}

	return prev
}
