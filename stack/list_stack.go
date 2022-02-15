package stack

import "fmt"

type LinkStack struct {
	Top *StackNode
}

type StackNode struct {
	Value int
	Next  *StackNode
}

func (s *LinkStack) Push(val int) bool {
	newNode := StackNode{Value: val}
	if s.Top == nil {
		s.Top = &newNode
	} else {
		newNode.Next = s.Top
		s.Top = &newNode
	}
	return true
}

func (s *LinkStack) Pop() int {
	if s.Top == nil {
		fmt.Println("stack is empty!")
		return -1
	}
	result := s.Top
	s.Top = result.Next
	return result.Value
}

func (s *LinkStack) Peek() int {
	if s.Top == nil {
		fmt.Println("stack is empty!")
		return -1
	}
	return s.Top.Value
}

func (s *LinkStack) IsEmpty() bool {
	return s.Top == nil
}

func (node *StackNode) Print() {
	if node == nil || node.Value == 0 {
		return
	} else {
		fmt.Print(node.Value, " ")
		node.Next.Print()
	}
}

func MainLinkStack() {
	linkStack := LinkStack{}
	linkStack.Top = &StackNode{Value: 10}
	linkStack.Top.Next = &StackNode{Value: 11}
	linkStack.Top.Next.Next = &StackNode{Value: 12}

	fmt.Println("-------- linkStack.Values  --------")
	linkStack.Top.Print()
	fmt.Println("")
	fmt.Println("-------- stack.IsEmpty()  --------")
	fmt.Println(linkStack.IsEmpty())
	fmt.Println("-------- stack.Pop()  --------")
	fmt.Println(linkStack.Pop())
	fmt.Println("-------- stack.Peek()  --------")
	fmt.Println(linkStack.Peek())
	fmt.Println("-------- linkStack.Values  --------")
	linkStack.Top.Print()
}
