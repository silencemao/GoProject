package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

func hasPathSum(root *tree.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func hasPathSum1(root *tree.TreeNode, sum int) bool {
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}
	if root.Left == nil && root.Right == nil && sum != 0 {
		return false
	}
	if root.Left != nil {
		if hasPathSum1(root.Left, sum-root.Left.Val) {
			return true
		}
	}
	if root.Right != nil {
		if hasPathSum1(root.Right, sum-root.Right.Val) {
			return true
		}
	}
	return false
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
	fmt.Println(hasPathSum(tRoot, 22))
	fmt.Println(hasPathSum1(tRoot, 22))
}
