/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    seen := make(map[*ListNode]struct{})

    for head != nil {
        if _, exists := seen[head]; exists {
            return true
        }
        seen[head] = struct{}{}
        head = head.Next
    }

    return false
}
