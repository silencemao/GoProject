package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

func help437(root *tree.TreeNode, targetSum int) int {
	res := 0
	if root == nil {
		return 0
	}
	if root.Val == targetSum {
		res += 1 // 此处不退出，沿着这条线继续找，若后续的子节点的值为0，则也是一种情况
	}
	res += help437(root.Left, targetSum-root.Val)
	res += help437(root.Right, targetSum-root.Val)
	return res
}

func pathSumIII(root *tree.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	res = help437(root, targetSum)
	res += pathSumIII(root.Left, targetSum)
	res += pathSumIII(root.Right, targetSum)
	return res
}

func main() {
	root := tree.Ints2Node([]int{10, 5, -3, 3, 2, tree.NilNode, 11, 3, -2, tree.NilNode, 1})
	fmt.Println(pathSumIII(root, 8))
}
