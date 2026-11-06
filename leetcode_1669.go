/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	r := b - a
	list1Prev := &ListNode{Next: list1}
	list1Curr := list1
	n := 0
	for list1Curr != nil {
		if n == a {
			list1Prev.Next = list2
			for range r {
				list1Curr = list1Curr.Next
			}
			list2Curr := list2
			for list2Curr != nil && list2Curr.Next != nil {
				list2Curr = list2Curr.Next
			}
			list2Curr.Next = list1Curr.Next
			break
		}
		list1Prev = list1Curr
		list1Curr = list1Curr.Next
		n++
	}

	return list1
}
