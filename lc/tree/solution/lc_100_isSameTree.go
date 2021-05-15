package main

import "GoProject/leetcode/tree"

func isSameTree(p, q *tree.TreeNode) bool {
	if p==nil && q==nil {
		return true
	}
	if p==nil || q==nil {
		return false
	}

	if p.Val==q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}

func main() {

}