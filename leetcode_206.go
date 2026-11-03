/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		// {nil <- 1 <- 2} = {nil, 1, 2}
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return prev
}
