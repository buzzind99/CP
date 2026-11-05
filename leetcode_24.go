/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	newHead := head.Next
	curr := head
	for curr != nil && curr.Next != nil {
		nextPair := curr.Next.Next
		currPair := curr.Next

		curr.Next = nextPair
		currPair.Next = curr

		if prev != nil {
			prev.Next = currPair
		}

		prev = curr
		curr = nextPair
	}

	return newHead
}
