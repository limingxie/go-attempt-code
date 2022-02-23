package graph

type ListGraph struct {
	Graphs map[int]*GraphNode
}

type GraphNode struct {
	Value int
	Next  *GraphNode
}

func (g *ListGraph) AddEdge(start, end int) {
	node := g.Graphs[start]
	if node == nil {
		g.Graphs[start] = &GraphNode{Value: end}
	} else {
		for node.Next != nil {
			node = node.Next
		}
		node.Next = &GraphNode{Value: end}
	}
}

func (g *ListGraph) RemoveEdge(start, end int) {
	node := g.Graphs[start]
	if node == nil {
		return
	} else {
		parentNode := node
		for node != nil && node.Value != end {
			parentNode = node
			node = node.Next
		}
		if node == nil {
			return
		} else {
			parentNode.Next = node.Next
		}
	}
}

func (g *ListGraph) DFSTraversal() {

}

func (g *ListGraph) BFSTraversal() {

}
