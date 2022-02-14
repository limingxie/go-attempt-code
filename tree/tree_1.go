package tree

import (
	"fmt"
	"math/rand"
	"time"
)

func Create(n int) *TreeNode {
	var t *TreeNode
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = Insert(t, temp)
	}
	return t
}

func Insert(t *TreeNode, v int) *TreeNode {
	if t == nil {
		return &TreeNode{v, nil, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = Insert(t.Left, v)
		return t
	}
	t.Right = Insert(t.Right, v)
	return t
}

func Traverse(t *TreeNode) {
	if t == nil {
		return
	}
	Traverse(t.Left)
	fmt.Println(t.Value, " ")
	Traverse(t.Right)
}
