package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
给定一个二叉树，我们在树的节点上安装摄像头。

节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。

计算监控树的所有节点所需的最小摄像头数量。

0 无覆盖
1 有摄像头
2 有覆盖

后续遍历

*/
func minCamera(root *tree.TreeNode, result *int) int {
	if root == nil {
		return 2
	}

	left := minCamera(root.Left, result)
	right := minCamera(root.Right, result)

	if left == 2 && right == 2 { // 左右孩子都有覆盖
		return 0 // 此时父节点是无覆盖，且暂时不需要立马加入一个摄像头，可以等该节点的父节点加入摄像头
	}
	if left == 0 || right == 0 { // 左孩子或右孩子是无覆盖，父节点立马加入摄像头，返回有摄像头
		*result += 1
		return 1
	}
	if left == 1 || right == 1 { // 左孩子或右孩子有摄像头，当前节点肯定被覆盖了，不需要加摄像头，直接返回有覆盖
		return 2
	}

	return 0
}

func minCameraCover(root *tree.TreeNode) int {
	var result int
	if minCamera(root, &result) == 0 {
		result += 1
	}
	return result
}

func main() {
	tRoot := tree.CreateTreeNode([]int{5, 4, 8, 11, tree.NilNode, 13, 4, 7, 2, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode, 1})
	result := minCameraCover(tRoot)
	fmt.Println(result)
}
