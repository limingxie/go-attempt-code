package linkedlist

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
	Size int
}

func (node *LinkedList) Insert(position int, number int) {
	if position > node.Size {
		return
	}

	newNode := &ListNode{Value: number}

	if position == 0 {
		newNode.Next = node.Head
		node.Head = newNode
		if node.Tail == nil {
			node.Tail = newNode
		}
	} else if position == node.Size {
		node.Tail.Next = newNode
		node.Tail = newNode
	} else {
		currentNode := node.Head
		for i := 0; i < position; i++ {
			currentNode = currentNode.Next
		}

		newNode.Next = currentNode.Next
		currentNode.Next = newNode
	}
	node.Size++
}

func (node *LinkedList) Delete1(number int) {
	if node.Head != nil && node.Head.Value == number {
		node.Head = node.Head.Next
		node.Size--
		if node.Size == 0 {
			node.Tail = node.Head
		}
	} else {
		prev := node.Head
		cur := node.Head
		for prev != nil && cur != nil {
			if cur.Value == number {
				if cur == node.Tail {
					node.Tail = prev
				}
				prev.Next = cur.Next
				node.Size--
				return
			}
			prev = cur
			cur = cur.Next
		}
	}
}

func (node *LinkedList) Delete(number int) {
	if node.Head == nil || node.Tail == nil {
		return
	}

	if node.Head.Value == number {
		node.Head = node.Head.Next
		node.Size--
		if node.Size == 0 {
			node.Tail = node.Head
		}
		return
	}

	preNode := node.Head
	curNode := node.Head
	for i := 0; i < node.Size; i++ {
		if curNode.Value == number {
			if curNode == node.Tail {
				node.Tail = preNode
			}
			preNode.Next = curNode.Next
			node.Size--
			return
		}
		preNode = curNode
		curNode = curNode.Next
	}
}

func (node *LinkedList) Update(old_number int, new_number int) int {
	currentNode := node.Head
	for i := 0; i < node.Size; i++ {
		if currentNode.Value == old_number {
			currentNode.Value = new_number
			return i
		}
		currentNode = currentNode.Next
	}
	return -1
}

func (node *LinkedList) Search(number int) int {
	currentNode := node.Head
	for i := 0; i < node.Size; i++ {
		if currentNode.Value == number {
			return i
		}
		currentNode = currentNode.Next
	}
	return -1
}

/*
func (node *LinkedList) Insert(position int, number int) {
	if position > node.Size {
		return
	}
	var newNode = &ListNode{Value: number}
	if position == 0 {
		newNode.Next = node.Head
		node.Head = newNode
		if node.Tail == nil {
			node.Tail = newNode
		}
	} else if position == node.Size {
		if node.Tail == nil {
			node.Tail = newNode
		} else {
			node.Tail.Next = newNode
			node.Tail = newNode
		}
	} else {
		var prev *ListNode = node.Head
		for i := 0; i < position; i++ {
			prev = prev.Next
		}
		next := prev.Next
		newNode.Next = next
		prev.Next = newNode
	}
	node.Size++
}

func (node *LinkedList) Delete(number int) {
	if node.Head != nil && node.Head.Value == number {
		node.Head = node.Head.Next
		node.Size--
		if node.Size == 0 {
			node.Tail = node.Head
		}
	} else {
		prev := node.Head
		cur := node.Head
		for prev != nil && cur != nil {
			if cur.Value == number {
				if cur == node.Tail {
					node.Tail = prev
				}
				prev.Next = cur.Next
				node.Size--
				return
			}
			prev = cur
			cur = cur.Next
		}
	}
}

func (node *LinkedList) Update(old_number int, new_number int) int {
	cur := node.Head
	for index := 0; cur != nil; index++ {
		if cur.Value == old_number {
			cur.Value = new_number
			return index
		}
		cur = cur.Next
	}
	return -1
}

func (node *LinkedList) Search(number int) int {
	cur := node.Head
	for index := 0; cur != nil; index++ {
		if cur.Value == number {
			return index
		}
		cur = cur.Next
	}
	return -1
}
*/
