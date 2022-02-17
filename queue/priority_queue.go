package queue

import "fmt"

type PriorityQueue struct {
	Head *Node
}

type Node struct {
	Value    string
	Priority int
	Next     *Node
}

func (p *PriorityQueue) Push(value string, priority int) {
	if p.Head == nil {
		p.Head = &Node{Value: value, Priority: priority}
		return
	}
	newNode := &Node{Value: value, Priority: priority}
	currentNode := p.Head
	for currentNode.Next != nil && currentNode.Next.Priority > priority {
		currentNode = currentNode.Next
	}
	newNode.Next = currentNode.Next
	currentNode.Next = newNode
}

func (p *PriorityQueue) Pop() *Node {
	if p.Head == nil {
		return nil
	}
	result := p.Head
	p.Head = p.Head.Next
	return result
}

func (p *PriorityQueue) Peek() *Node {
	return p.Head
}

func (p *PriorityQueue) IsEmpty() bool {
	return p.Head == nil
}

func (n *Node) Print() {
	if n == nil {
		return
	}
	fmt.Print(n.Value, " ")
	n.Next.Print()
}

func MainPriorityQueue() {
	priorityQueue := PriorityQueue{}
	priorityQueue.Push("a", 100)
	priorityQueue.Push("b", 83)
	priorityQueue.Push("c", 64)
	priorityQueue.Push("e", 37)
	priorityQueue.Push("f", 23)

	fmt.Println("-------- Print  --------")
	priorityQueue.Head.Print()
	fmt.Println("")

	fmt.Println("-------- Pop  --------")
	fmt.Printf("%+v", priorityQueue.Pop())
	fmt.Println("")
	priorityQueue.Head.Print()
	fmt.Println("")

	fmt.Println("-------- Push z --------")
	priorityQueue.Push("z", 53)
	priorityQueue.Head.Print()
	fmt.Println("")

	fmt.Println("-------- Peek()  --------")
	fmt.Printf("%+v", priorityQueue.Peek())
}
