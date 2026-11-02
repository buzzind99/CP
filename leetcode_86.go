/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	lessDummy := &ListNode{}
	greaterDummy := &ListNode{}
	less, greater := lessDummy, greaterDummy

	for curr := head; curr != nil; {
		if curr.Val < x {
			less.Next = curr
			less = less.Next
		} else {
			greater.Next = curr
			greater = greater.Next
		}
		curr = curr.Next
	}
	greater.Next = nil
	less.Next = greaterDummy.Next

	return lessDummy.Next
}
