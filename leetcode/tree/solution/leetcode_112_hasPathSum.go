package main

import (
	"fmt"
	"myGoProject/leetcode/tree"
)

func hasPathSum(root *tree.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
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
}
