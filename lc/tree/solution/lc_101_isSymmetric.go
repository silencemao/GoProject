package main

import "GoProject/leetcode/tree"


func symmetricHelper(l, r *tree.TreeNode) bool {
	if l==nil && r==nil {
		return true
	}
	if l==nil || r==nil {
		return false
	}
	if l.Val==r.Val {
		return symmetricHelper(l.Left, r.Right) && symmetricHelper(l.Right, r.Left)
	}
	return false
}

func isSymmetric(root *tree.TreeNode) bool {
	if root==nil {
		return true
	}
	return symmetricHelper(root.Left, root.Right)
}

func main() {

}
