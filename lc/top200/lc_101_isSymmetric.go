package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

func help101(l, r *tree.TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	return help101(l.Left, r.Right) && help101(l.Right, r.Left)
}
func isSymmetric(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	return help101(root.Left, root.Right)
}

func main() {
	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
			},
			Right: &tree.TreeNode{
				Val: 5,
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 5,
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
	}
	//fmt.Println(levelOrderTraversal1(root))
	fmt.Println(isSymmetric(root))
}
