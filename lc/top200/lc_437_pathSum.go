package main

import "GoProject/leetcode/tree"

func findPath(root *tree.TreeNode, targetSum int) int {
	res := 0
	if root == nil {
		return 0
	}
	if root.Val == targetSum {
		res += 1
	}
	res += findPath(root.Left, targetSum-root.Val)
	res += findPath(root.Right, targetSum-root.Val)
	return res
}

func pathSum(root *tree.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := findPath(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}
