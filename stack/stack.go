package stack

import "fmt"

type Stack struct {
	Capacity int
	Top      int
	Values   []int
}

func (s *Stack) Push(val int) bool {
	if s.Top >= s.Capacity {
		fmt.Println("Stack Overflow!")
		return false
	}
	s.Top++
	s.Values[s.Top] = val
	return true
}

func (s *Stack) Pop() int {
	if s.Top < 0 {
		fmt.Println("Stack Underflow!")
		return 0
	}
	result := s.Values[s.Top]
	s.Top--
	return result
}

func (s *Stack) Peek() int {
	if s.Top < 0 {
		fmt.Println("Stack Underflow!")
		return 0
	}
	result := s.Values[s.Top]
	return result
}

func (s *Stack) IsEmpty() bool {
	return s.Top < 0
}

func MainStack() {
	stack := Stack{}
	stack.Capacity = 10
	stack.Top = -1
	stack.Values = make([]int, stack.Capacity)

	stack.Push(10)
	stack.Push(11)
	stack.Push(12)
	fmt.Println("-------- stack.Values  --------")
	fmt.Printf("%+v", stack.Values)
	fmt.Println("")
	fmt.Println("-------- stack.IsEmpty()  --------")
	fmt.Println(stack.IsEmpty())
	fmt.Println("-------- stack.Pop()  --------")
	fmt.Println(stack.Pop())
	fmt.Println("-------- stack.Peek()  --------")
	fmt.Println(stack.Peek())
}
