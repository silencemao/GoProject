package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。

*/
func help113(res *[][]int, tmp []int, root *tree.TreeNode, targetSum int) {
	if root == nil {
		return
	}
	tmp = append(tmp, root.Val)
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		*res = append(*res, append([]int{}, tmp...))
		return
	}
	help113(res, tmp, root.Left, targetSum)
	help113(res, tmp, root.Right, targetSum)
	tmp = tmp[:len(tmp)-1] // 一次撤销操作 从root位置开始，先执行left,向下执行到叶节点，回退到root时，先不弹出root值，继续执行root的right位置
}

func pathSum(root *tree.TreeNode, targetSum int) [][]int {
	var res [][]int
	help113(&res, []int{}, root, targetSum)
	fmt.Println(res)
	return res
}

/*

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
*/
func main() {
	tRoot := tree.CreateTreeNode([]int{5, 4, 8, 11, tree.NilNode, 13, 4, 7, 2, tree.NilNode, 1})
	targetSum := 22
	fmt.Println(pathSum(tRoot, targetSum))
}
