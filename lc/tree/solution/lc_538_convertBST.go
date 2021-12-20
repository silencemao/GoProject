package main

import "GoProject/leetcode/tree"

/*
把二叉搜索树转换为累加树
给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），
使每个节点 node的新值等于原树中大于或等于 node.val的值之和。

提醒一下，二叉搜索树满足下列约束条件：
节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。
*/
func help538(root *tree.TreeNode, sum *int) {
	if root == nil {
		return
	}
	help538(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	help538(root.Left, sum)
}

func convertBST(root *tree.TreeNode) *tree.TreeNode {
	sum := 0
	help538(root, &sum)
	return root
}

func main() {
	tRoot := tree.CreateTreeNode([]int{4, 2, 7, 1, 3, tree.NilNode, tree.NilNode})
	tRoot = convertBST(tRoot)
}
