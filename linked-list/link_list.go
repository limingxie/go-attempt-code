package linkedlist

import "fmt"

type LinkList struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (l *LinkList) ReverseLinkedList() {
	if l == nil || l.Head == nil {
		return
	}
	reverseLink := LinkList{}
	node := l.Head
	reverseLink.Head = &Node{Value: node.Value}
	for node.Next != nil {
		newNode := reverseLink.Head
		reverseLink.Head = &Node{Value: node.Next.Value}
		reverseLink.Head.Next = newNode
		node = node.Next
	}
	l.Head = reverseLink.Head
}

func (node *Node) Print() {
	if node == nil || node.Value == 0 {
		return
	} else {
		fmt.Print(node.Value, " ")
		node.Next.Print()
	}
}

func MainFuncReverseLinkedList() {
	linkList := LinkList{}
	linkList.Head = &Node{Value: 1}
	linkList.Head.Next = &Node{Value: 2}
	linkList.Head.Next.Next = &Node{Value: 3}
	linkList.Head.Next.Next.Next = &Node{Value: 4}

	linkList.ReverseLinkedList()

	linkList.Head.Print()
}
