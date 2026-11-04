/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	modified := head.Next
	curr := modified
	for curr != nil {
		sum := 0
		for curr.Val != 0 {
			sum += curr.Val
			curr = curr.Next
		}

		modified.Val = sum
		curr = curr.Next

		modified.Next = curr
		modified = modified.Next
	}

	return head.Next
}
