type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{Val: -1, Next: nil}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	curr := this.Next
	for range index {
		curr = curr.Next
	}

	if curr == nil {
		return -1
	}

	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &MyLinkedList{Val: val, Next: this.Next}
	this.Next = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &MyLinkedList{Val: val}
	curr := this
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	}

	curr := this
	for range index {
		curr = curr.Next
	}
	newNode := &MyLinkedList{Val: val, Next: curr.Next}
	curr.Next = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		index = 0
	}

	curr := this
	for range index {
		curr = curr.Next
	}

	if curr == nil || curr.Next == nil {
		return
	}
	curr.Next = curr.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
