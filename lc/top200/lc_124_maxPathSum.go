package main

import (
	"GoProject/leetcode/tree"
	"GoProject/leetcode/util"
	"fmt"
)

/*
二叉树最大路径和

路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。
该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。

输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6

*/
func help124(root *tree.TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	leftMax := util.MaxInt(help124(root.Left, res), 0)
	rightMax := util.MaxInt(help124(root.Right, res), 0)

	cur := root.Val + leftMax + rightMax
	*res = util.MaxInt(*res, cur)

	return root.Val + util.MaxInt(leftMax, rightMax)
}

func maxPathSum(root *tree.TreeNode) int {
	res := -1<<31 - 1
	fmt.Println(help124(root, &res))
	return res
}
func main() {
	root := &tree.TreeNode{Val: 5, Left: &tree.TreeNode{Val: 4, Left: &tree.TreeNode{Val: 1}}, Right: &tree.TreeNode{Val: 2}}
	tree.LevelOrder(root)
	fmt.Println(maxPathSum(root))
}
