package linkedlist

import "fmt"

type CircleLinkedList struct {
	Head *ListNode
	Tail *ListNode
	Size int
}

//约瑟夫问题
func (c *CircleLinkedList) JosephusCircleLinkedList(startNo, count, sum int) {
	for i := 1; i < startNo; i++ {
		c.Head = c.Head.Next
		c.Tail = c.Tail.Next
	}

	for {
		if c.Head == c.Tail {
			break
		}
		for i := 0; i < count-1; i++ {
			c.Head = c.Head.Next
			c.Tail = c.Tail.Next
		}
		fmt.Printf("出队元素：%d\n", c.Head.Value)
		c.Head = c.Head.Next
		c.Tail.Next = c.Head

	}
	fmt.Printf("最后剩下的元素是：%d", c.Head.Value)
}

func MainJosephusCircleLinkedList() {
	c := CircleLinkedList{}
	c.Head = &ListNode{Value: 1}
	c.Head.Next = &ListNode{Value: 2}
	c.Head.Next.Next = &ListNode{Value: 3}
	c.Head.Next.Next.Next = &ListNode{Value: 4}
	c.Head.Next.Next.Next.Next = &ListNode{Value: 5}
	c.Tail = c.Head.Next.Next.Next.Next
	c.Tail.Next = c.Head
	c.Size = 5
	c.JosephusCircleLinkedList(1, 2, 5)
}
