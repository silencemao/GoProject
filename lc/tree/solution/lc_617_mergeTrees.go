package main

import "GoProject/leetcode/tree"

func mergeTrees(root1 *tree.TreeNode, root2 *tree.TreeNode) *tree.TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return &tree.TreeNode{Val: root2.Val}
	}
	if root2 == nil {
		return &tree.TreeNode{Val: root1.Val}
	}

	root := &tree.TreeNode{Val: root1.Val + root2.Val}
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}

func main() {

}
