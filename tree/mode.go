package tree

import "fmt"

type BinaryTree struct {
	Value    int
	LeftNode *BinaryTree
	RighNode *BinaryTree
}

func (binaryTree *BinaryTree) Print() {
	fmt.Print(binaryTree.Value, " ")
}

func (binaryTree *BinaryTree) PreNode() {
	if binaryTree == nil {
		return
	}
	binaryTree.Print()
	binaryTree.LeftNode.Print()
	binaryTree.RighNode.Print()
}

//创建二叉树
func createBinaryTree(i int, nums []int) *BinaryTree {
	tree := &BinaryTree{nums[i], nil, nil}
	//左节点的数组下标为1,3,5...2*i+1
	if i < len(nums) && 2*i+1 < len(nums) {
		tree.LeftNode = createBinaryTree(2*i+1, nums)
	}
	//右节点的数组下标为2,4,6...2*i+2
	if i < len(nums) && 2*i+2 < len(nums) {
		tree.RighNode = createBinaryTree(2*i+2, nums)
	}
	return tree
}
