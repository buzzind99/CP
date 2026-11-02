/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sz := 1
	for curr := head; curr != nil; curr = curr.Next {
		sz++
	}

	pos := 1
	dummy := &ListNode{Next: head}
	prev := dummy
	for curr := head; curr != nil; curr = curr.Next {
		if pos == sz-n {
			prev.Next = curr.Next
			curr.Next = nil
			break
		}
		prev = curr
		prev.Next = curr.Next
		pos++
	}

	return dummy.Next
}
