/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	for range left - 1 {
		prev = prev.Next
	}

	curr := prev.Next
	next := curr.Next
	for range right - left {
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = curr.Next
	}

	return dummy.Next
}
