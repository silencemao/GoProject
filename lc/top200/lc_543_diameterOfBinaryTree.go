package main

import (
	"GoProject/leetcode/tree"
	"GoProject/leetcode/util"
)

/*
二叉树的最大路径
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
两结点之间的路径长度是以它们之间边的数目表示

最大路径不一定能经过根节点，若二叉树的右子树为空，则最大路径可能是左子树的某一个节点下的路径
计算某个节点下左子树的最大深度l，右子树的最大深度r，直径=l+r+1
*/
func depthHelp543(root *tree.TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := depthHelp543(root.Left, res)
	r := depthHelp543(root.Right, res)
	*res = util.MaxInt(1+l+r, *res)
	return util.MaxInt(l, r) + 1
}

func diameterOfBinaryTree(root *tree.TreeNode) int {
	res := 1
	if root != nil {
		depthHelp543(root, &res)
	}
	return res - 1 // 节点之间边的个数
}

func main() {

}
