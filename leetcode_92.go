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

	start := prev.Next
	after := start.Next

	for range right - left {
		start.Next = after.Next
		after.Next = prev.Next
		prev.Next = after
		after = start.Next
	}

	return dummy.Next
}
