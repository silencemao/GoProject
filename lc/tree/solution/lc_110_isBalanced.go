package main

import (
	"fmt"
	"math"
	"myGoProject/leetcode/tree"
)

func dfsHelper(tRoot *tree.TreeNode) int {
	if tRoot==nil{
		return 0
	}
	mLeft := dfsHelper(tRoot.Left)
	if mLeft==-1 {
		return -1
	}
	mRight := dfsHelper(tRoot.Right)
	if mRight==-1 {
		return -1
	}
	if math.Abs(float64(mLeft-mRight)) > 1 {
		return -1
	}
	return int(math.Max(float64(mLeft), float64(mRight)) + 1)
}

func isBalanced(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	mLeft := tree.MaxDepth(root.Left)
	mRight := tree.MaxDepth(root.Right)
	if math.Abs(float64(mLeft-mRight)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func isBalanced1(root *tree.TreeNode) bool {
	return dfsHelper(root) != -1
}

func main() {
	tRoot := tree.CreateTreeNode([]int{1,2,2,3,3, tree.NilNode, tree.NilNode,4,4})
	fmt.Println(isBalanced(tRoot))
	fmt.Println(isBalanced1(tRoot))
}
