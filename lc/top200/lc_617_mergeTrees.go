package main

import "GoProject/leetcode/tree"

/*

合并二叉树 前序遍历
*/
func mergeTrees(root1 *tree.TreeNode, root2 *tree.TreeNode) *tree.TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func main() {
	r1 := &tree.TreeNode{Val: 1, Left: &tree.TreeNode{Val: 3, Left: &tree.TreeNode{Val: 5}}, Right: &tree.TreeNode{Val: 2}}
	r2 := &tree.TreeNode{Val: 2, Left: &tree.TreeNode{Val: 1, Right: &tree.TreeNode{Val: 4}}, Right: &tree.TreeNode{Val: 3, Right: &tree.TreeNode{Val: 7}}}

	r := mergeTrees(r1, r2)
	tree.LevelOrder(r)

}
