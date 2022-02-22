package tree

type Heap struct {
	Head *HeapNode
}

type HeapNode struct {
	Value int
	Next  *HeapNode
}

func (h *HeapNode) Add(val int) {

}
func (h *HeapNode) Poll() *HeapNode {
	return nil
}

func (h *HeapNode) HeapifyUp() {

}

func (h *HeapNode) HeapifyDown() {

}

func (h *HeapNode) Peek() *HeapNode {
	return nil
}
