package linkedlist

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func (node *ListNode) Print() {
	if node == nil || node.Value == 0 {
		return
	} else {
		fmt.Print(node.Value, " ")
		node.Next.Print()
	}
}

func Mainfunc() {
	var linkedlist LinkedList
	linkedlist.Insert(0, 1)
	linkedlist.Insert(1, 2)
	linkedlist.Insert(2, 3)
	linkedlist.Insert(3, 4)
	linkedlist.Insert(4, 5)
	linkedlist.Insert(5, 6)
	linkedlist.Insert(6, 7)
	linkedlist.Insert(7, 8)
	fmt.Println("----insert 1,2,3,4,5,6,7,8----")
	linkedlist.Head.Print()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("------update 5 => 55 ------")
	linkedlist.Update(5, 55)
	linkedlist.Head.Print()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("-------- delete 55 --------")
	linkedlist.Delete(55)
	linkedlist.Head.Print()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("-------- delete 9 --------")
	linkedlist.Delete(9)
	linkedlist.Head.Print()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("-------- search 4  --------")
	fmt.Println("index:", linkedlist.Search(4))
}
