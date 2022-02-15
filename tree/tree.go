package tree

import "fmt"

type BTSTree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (b *BTSTree) Get(val int) *TreeNode {
	currentNode := b.Root
	for currentNode != nil && currentNode.Value != val {
		if currentNode.Value == val {
			break
		} else if currentNode.Value > val {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	return currentNode
}

func (b *BTSTree) Insert(val int) {
	if b.Root == nil {
		b.Root = &TreeNode{Value: val}
		return
	}
	currentNode := b.Root
	parentNode := currentNode
	for currentNode != nil {
		parentNode = currentNode
		if currentNode.Value > val {
			currentNode = currentNode.Left
			if currentNode == nil {
				parentNode.Left = &TreeNode{Value: val}
				return
			}
		} else if currentNode.Value < val {
			currentNode = currentNode.Right
			if currentNode == nil {
				parentNode.Right = &TreeNode{Value: val}
				return
			}
		} else {
			return // currentNode.Value == val
		}
	}
}

func (b *BTSTree) Delete(val int) bool {
	if b.Root == nil {
		fmt.Println("tree is empty!")
		return false
	}
	currentNode := b.Root
	parentNode := currentNode
	isLeftChild := false

	for currentNode != nil && currentNode.Value != val {
		parentNode = currentNode
		if currentNode.Value == val {
			break
		} else if currentNode.Value > val {
			currentNode = currentNode.Left
			isLeftChild = true
		} else {
			currentNode = currentNode.Right
			isLeftChild = false
		}
	}

	if currentNode.Left == nil && currentNode.Right == nil {
		if b.Root == currentNode {
			b.Root = nil
		} else if isLeftChild {
			parentNode.Left = nil
		} else {
			parentNode.Right = nil
		}
	}

	if currentNode.Left != nil && currentNode.Right == nil {
		if b.Root == currentNode {
			b.Root = currentNode.Left
		} else if isLeftChild {
			parentNode.Left = currentNode.Left
		} else {
			parentNode.Right = currentNode.Left
		}
	} else if currentNode.Left == nil && currentNode.Right != nil {
		if b.Root == currentNode {
			b.Root = currentNode.Right
		} else if isLeftChild {
			parentNode.Left = currentNode.Right
		} else {
			parentNode.Right = currentNode.Right
		}
	}

	if currentNode.Left != nil && currentNode.Right != nil {
		treeNode := currentNode.GetSubTreeForDelete()
		if b.Root == currentNode {
			b.Root = treeNode
		} else if isLeftChild {
			parentNode.Left = treeNode
		} else {
			parentNode.Right = treeNode
		}
	}

	return true
}

func (node *TreeNode) GetSubTreeForDelete() *TreeNode {
	parentNode := node
	currentNode := node.Right
	for currentNode != nil {
		if currentNode.Left == nil {
			break
		}
		parentNode = currentNode
		currentNode = currentNode.Left

	}
	if currentNode == node.Right {
		parentNode.Right = currentNode.Right
		return parentNode
	} else {
		parentNode.Left = currentNode.Right
		currentNode.Left = node.Left
		currentNode.Right = node.Right
		return currentNode
	}
}

func (node *TreeNode) PreOrderTraversalPrint() {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	node.Left.PreOrderTraversalPrint()
	node.Right.PreOrderTraversalPrint()
}

func (node *TreeNode) InOrderTraversalPrint() {
	if node == nil {
		return
	}
	node.Left.PreOrderTraversalPrint()
	fmt.Println(node.Value)
	node.Right.PreOrderTraversalPrint()
}

func (node *TreeNode) PostOrderTraversalPrint() {
	if node == nil {
		return
	}
	node.Left.PreOrderTraversalPrint()
	node.Right.PreOrderTraversalPrint()
	fmt.Println(node.Value)
}

func MainBTSTree() {
	fmt.Println("aa")
}
