/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	seen := make(map[*ListNode]struct{})
	curr := headA
	for curr != nil {
		seen[curr] = struct{}{}
		curr = curr.Next
	}
	intersect := headB
	for intersect != nil {
		if _, exists := seen[intersect]; exists {
			return intersect
		}
		intersect = intersect.Next
	}

	return nil
}
