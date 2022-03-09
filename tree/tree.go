package tree

import "fmt"

type BSTTree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (b *BSTTree) Get(val int) *TreeNode {
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

func (b *BSTTree) GetParent(val int) *TreeNode {
	if b.Root == nil {
		return nil
	}
	return b.Root.GetParent(val)
}

func (t *TreeNode) GetParent(val int) *TreeNode {
	if t == nil {
		return nil
	}
	if (t.Left != nil && t.Left.Value == val) || (t.Right != nil && t.Right.Value == val) {
		return t
	}
	if t.Value > val && t.Left != nil {
		return t.Left.GetParent(val)
	} else if t.Value < val && t.Right != nil {
		return t.Right.GetParent(val)
	}
	return nil
}

func (b *BSTTree) Insert_recursive(val int) {
	if b.Root == nil {
		b.Root = &TreeNode{Value: val}
		return
	}
	b.Root = Insert_recursive(b.Root, val)
}

func Insert_recursive(t *TreeNode, val int) *TreeNode {
	if t == nil {
		t = &TreeNode{Value: val}
		return t
	}

	if t.Value > val {
		t.Left = Insert_recursive(t.Left, val)
	} else if t.Value < val {
		t.Right = Insert_recursive(t.Right, val)
	}
	return t

}

func (b *BSTTree) Insert(val int) {
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

func (b *BSTTree) Delete(val int) bool {
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
	} else if currentNode.Left != nil && currentNode.Right == nil {
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
	} else if currentNode.Left != nil && currentNode.Right != nil {
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
	fmt.Print(node.Value, " ")
	node.Left.PreOrderTraversalPrint()
	node.Right.PreOrderTraversalPrint()
}

func (node *TreeNode) InOrderTraversalPrint() {
	if node == nil {
		return
	}
	node.Left.InOrderTraversalPrint()
	fmt.Print(node.Value, " ")
	node.Right.InOrderTraversalPrint()
}

func (node *TreeNode) PostOrderTraversalPrint() {
	if node == nil {
		return
	}
	node.Left.PostOrderTraversalPrint()
	node.Right.PostOrderTraversalPrint()
	fmt.Print(node.Value, " ")
}

func MainBSTTree() {
	btsTree := BSTTree{Root: &TreeNode{Value: 57}}
	btsTree.Insert(21)
	btsTree.Insert(88)
	btsTree.Insert(12)
	btsTree.Insert(36)
	btsTree.Insert(69)
	btsTree.Insert(97)
	btsTree.Insert(7)
	btsTree.Insert(14)
	btsTree.Insert(24)
	btsTree.Insert(47)
	btsTree.Insert(61)
	btsTree.Insert(73)
	btsTree.Insert(92)
	btsTree.Insert(99)

	fmt.Println("------------ Print --------------")
	btsTree.Root.PreOrderTraversalPrint()
	fmt.Println("")
	btsTree.Root.PostOrderTraversalPrint()
	fmt.Println("")
	btsTree.Root.InOrderTraversalPrint()
	fmt.Println("")

	fmt.Println("------------ Get 24 --------------")
	fmt.Printf("%+v\n", btsTree.Get(24))
	fmt.Printf("%+v\n", btsTree.GetParent(24))

	fmt.Println("----------- Insert 95 -------------")
	btsTree.Insert(95)
	btsTree.Root.PreOrderTraversalPrint()
	fmt.Println("")
	btsTree.Root.PostOrderTraversalPrint()
	fmt.Println("")
	btsTree.Root.InOrderTraversalPrint()
	fmt.Println("")

	fmt.Println("----------- Delete 88 -------------")
	btsTree.Delete(88)
	btsTree.Root.PreOrderTraversalPrint()
	fmt.Println("")
	btsTree.Root.PostOrderTraversalPrint()
	fmt.Println("")
	btsTree.Root.InOrderTraversalPrint()
	fmt.Println("")
}
