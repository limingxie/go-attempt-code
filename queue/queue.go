package queue

import "fmt"

type Queue struct {
	Front    int
	Rear     int
	Size     int
	Capacity int
	Values   []int
}

func (q *Queue) Enqueue(item int) {
	if q.IsFull() {
		fmt.Println("queue is full!")
		return
	}

	q.Values[q.Rear] = item
	q.Rear = (q.Rear + 1) % q.Capacity
	q.Size++
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("queue is empty!")
		return -1
	}
	front := q.Values[q.Front]
	q.Front = (q.Front + 1) % q.Capacity
	return front
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		fmt.Println("queue is empty!")
		return -1
	}

	return q.Values[q.Front]
}

func (q *Queue) IsFull() bool {
	return q.Size == q.Capacity
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}

func Mainfunc() {
	queue := Queue{Capacity: 10}
	queue.Values = make([]int, queue.Capacity)
	queue.Enqueue(11)
	queue.Enqueue(12)
	queue.Enqueue(13)
	fmt.Println("-------- queue.Values  --------")
	fmt.Printf("%+v", queue.Values)
	fmt.Println("")
	fmt.Println("-------- Front, Rear  --------")
	fmt.Printf("Front: %v, Rear: %v", queue.Front, queue.Rear)
	fmt.Println("")
	queue.Enqueue(14)
	fmt.Println("-------- queue.Values  --------")
	fmt.Printf("%+v", queue.Values)
	fmt.Println("")
	fmt.Println("-------- Front, Rear  --------")
	fmt.Printf("Front: %v, Rear: %v", queue.Front, queue.Rear)
	fmt.Println("")
	fmt.Println("-------- queue.IsEmpty  --------")
	fmt.Println(queue.IsEmpty())
	fmt.Println("-------- queue.IsFull  --------")
	fmt.Println(queue.IsFull())
	fmt.Println("-------- queue.Dequeue  --------")
	fmt.Println(queue.Dequeue())
	fmt.Println("-------- Front, Rear  --------")
	fmt.Printf("Front: %v, Rear: %v", queue.Front, queue.Rear)
}
