package main

import (
	"fmt"
	"myGoProject/leetcode/tree"
)

func sumOfLeftLeaves(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left==nil && root.Left.Right==nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
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
	tRoot := tree.CreateTreeNode([]int{5, 4, 8, 11, tree.NilNode, 13, 4, 7, 2, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode, 1})
	fmt.Println(sumOfLeftLeaves(tRoot))
}
