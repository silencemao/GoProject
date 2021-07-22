package main

import (
	"GoProject/leetcode/tree"
	"GoProject/leetcode/util"
)

/*
二叉树的最大路径
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

*/
func depthHelp543(root *tree.TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := depthHelp543(root.Left, res)
	r := depthHelp543(root.Right, res)
	*res = 1 + l + r
	return util.MaxInt(l, r) + 1
}

func diameterOfBinaryTree(root *tree.TreeNode) int {
	res := 1
	if root != nil {
		depthHelp543(root, &res)
	}
	return res
}

func main() {

}
