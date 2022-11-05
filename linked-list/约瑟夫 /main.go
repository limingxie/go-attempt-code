package main

import (
	"fmt"
)

type CircleLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	Value    int
	NextNode *Node
}

func (c *CircleLinkedList) JosephusCircleLinkedList(startNo, count int) {
	node := c.Head
	endNode := c.Head
	for i := 1; i < startNo; i++ {
		node = node.NextNode
	}
	for i := 1; i <= (startNo+c.Size-2)%c.Size; i++ {
		endNode = endNode.NextNode
	}
	fmt.Println("-----------当前node和尾随node------------")
	fmt.Printf("%p => %v \n", node, node)
	fmt.Printf("%p => %v \n", endNode, endNode)
	for {
		for i := 1; i <= count; i++ {
			node = node.NextNode
			endNode = endNode.NextNode
		}

		endNode.NextNode = node.NextNode
		fmt.Printf("出列 %v\n", node.Value)
		node = node.NextNode
		if node == endNode {
			fmt.Printf("出列 %v\n", node.Value)
			break
		}
	}
}

func CreateCircleLinkedList(count int) *CircleLinkedList {
	c := &CircleLinkedList{Size: count}
	node := &Node{Value: 1}
	c.Head = node

	for i := 1; i < count; i++ {
		n := &Node{Value: i + 1}
		node.NextNode = n
		node = n
	}
	node.NextNode = c.Head
	return c
}

func (c *CircleLinkedList) Print() {
	node := c.Head
	for i := 0; i < c.Size; i++ {
		fmt.Printf("%p => %v \n", node, node)
		node = node.NextNode
	}
}

func main() {
	c := CreateCircleLinkedList(9)
	c.Print()

	c.JosephusCircleLinkedList(2, 2)
}
