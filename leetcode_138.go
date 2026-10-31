/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	curr := head
	for curr != nil {
		copy := &Node{Val: curr.Val, Next: curr.Next}
		curr.Next = copy
		curr = copy.Next
	}

	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	dummy := &Node{}
	copied := dummy
	curr = head
	for curr != nil {
		copy := curr.Next
		curr.Next = copy.Next
		copied.Next = copy
		copied = copy
		curr = curr.Next
	}

	return dummy.Next
}
