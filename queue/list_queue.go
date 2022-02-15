package queue

import "fmt"

type ListQueue struct {
	Front *QueueNode
	Rear  *QueueNode
}

type QueueNode struct {
	Value int
	Next  *QueueNode
}

func (q *ListQueue) Enqueue(val int) {
	newNode := QueueNode{Value: val}
	if q.Rear == nil {
		q.Front = &newNode
		q.Rear = &newNode
		return
	}

	q.Rear.Next = &newNode
	q.Rear = &newNode
}

func (q *ListQueue) Dequeue() int {
	if q.Front == nil {
		fmt.Println("queue is empty!")
		return -1
	}
	front := q.Front
	q.Front = q.Front.Next
	if q.Front == nil {
		q.Rear = nil
	}

	return front.Value
}

func (node *QueueNode) Print() {
	if node == nil || node.Value == 0 {
		return
	} else {
		fmt.Print(node.Value, " ")
		node.Next.Print()
	}
}
