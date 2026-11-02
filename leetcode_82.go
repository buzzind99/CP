/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for curr := head; curr != nil; {
		next := curr.Next
		if next != nil && curr.Val == next.Val {
			for next != nil && next.Val == curr.Val {
				next = next.Next
			}
			prev.Next = next
			curr = next
			continue
		}
		prev = curr
		curr = next
	}

	return dummy.Next
}
