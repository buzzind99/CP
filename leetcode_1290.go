/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	sum, pow, count := 0, 0, 0
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		pow++
		if curr.Val == 1 {
			sum += 1 << (count - pow)
		}
		curr = curr.Next
	}

	return sum
}
