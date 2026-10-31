/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	copyMap := make(map[*Node]*Node)
	curr := head
	for curr != nil {
		copyMap[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		if curr.Next != nil {
			copyMap[curr].Next = copyMap[curr.Next]
		}
		if curr.Random != nil {
			copyMap[curr].Random = copyMap[curr.Random]
		}
		curr = curr.Next
	}

	return copyMap[head]
}
