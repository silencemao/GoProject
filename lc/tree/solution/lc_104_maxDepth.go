package main

import (
	"GoProject/leetcode/tree"
	"GoProject/leetcode/util"
	"fmt"
)

func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	lD := maxDepth(root.Left)
	fmt.Println("l", lD)
	rD := maxDepth(root.Right)
	fmt.Println("r", rD)
	return 1 + util.MaxInt(lD, rD)
	//return int(1 + math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}

/*
n叉树的最大深度
*/
type Node struct {
	Val      int
	Children []*Node
}

func maxNTreeDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	for i := 0; i < len(root.Children); i++ {
		depth = max(depth, maxNTreeDepth(root.Children[i]))
	}
	return depth + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
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
	fmt.Println(maxDepth(tRoot))
}
