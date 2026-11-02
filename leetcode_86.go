/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	anchor := &ListNode{Next: head}
	prev := anchor
	lastAnchor := 0
	pos := 1
	curr := head
	for curr != nil {
		if curr.Val < x {
			if pos < x && head.Val < x && lastAnchor == pos-1 {
				anchor = curr
				lastAnchor++
			} else {
				prev.Next = curr.Next
				curr.Next = anchor.Next
				if anchor.Next == head {
					head = curr
				}
				anchor.Next = curr
				anchor = curr
				lastAnchor++
			}
		}
		prev = curr
		curr = curr.Next
		pos++
	}

	return head
}
