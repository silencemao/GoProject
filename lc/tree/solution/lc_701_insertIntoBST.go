package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
给定一个数字 插入二叉搜索树中
*/
func help701(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return &tree.TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = help701(root.Right, val)
	}
	if root.Val > val {
		root.Left = help701(root.Left, val)
	}
	return root
}

/*
       4
    2     7
  1   3
*/

func insertIntoBST(root *tree.TreeNode, val int) *tree.TreeNode {
	return help701(root, val)
}
func main() {
	tRoot := tree.CreateTreeNode([]int{4, 2, 7, 1, 3, tree.NilNode, tree.NilNode})
	root := insertIntoBST(tRoot, 5)
	fmt.Println(root.Val)
}
