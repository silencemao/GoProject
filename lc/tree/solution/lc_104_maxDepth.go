package main

import (
	"math"
	"GoProject/leetcode/tree"
)

func maxDepth(root *tree.TreeNode) int {
	if root==nil {
		return 0
	}
	return int(1 + math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}

func main() {

}