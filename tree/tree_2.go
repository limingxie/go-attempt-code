package tree

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (binaryTree *TreeNode) Print() {
	fmt.Print(binaryTree.Value, " ")
}

// 增加二叉树左节点
func (node *TreeNode) AddLeftNode(value int) {
	children := TreeNode{Value: value}
	node.Left = &children
}

// 增加二叉树右节点
func (node *TreeNode) AddRightNode(value int) {
	children := TreeNode{Value: value}
	node.Right = &children
}

func (node *TreeNode) FrontNode() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.FrontNode()
	node.Right.FrontNode()
}

func (node *TreeNode) MiddleNode() {
	if node == nil {
		return
	}
	node.Left.MiddleNode()
	node.Print()
	node.Right.MiddleNode()
}

func (node *TreeNode) BehindNode() {
	if node == nil {
		return
	}
	node.Left.BehindNode()
	node.Right.BehindNode()
	node.Print()
}

func main() {

}
