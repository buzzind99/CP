/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	seen := make(map[*ListNode]struct{})
	curr := head
	for curr != nil {
		if _, exists := seen[curr]; exists {
			break
		}
		seen[curr] = struct{}{}
		curr = curr.Next
	}

	return curr
}
