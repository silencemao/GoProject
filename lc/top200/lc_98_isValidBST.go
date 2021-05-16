package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
平衡二叉搜索树
左子树节点的值全部小于根结点的值
右子树节点的值全部小于根结点的值
左右子树分别也是二叉搜索树
*/
var pre int = -1 << 63

// 中序遍历
func isValidBST1(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBST1(root.Left) {
		return false
	}
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return isValidBST1(root.Right)
}

func isBST(root *tree.TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return isBST(root.Left, min, root.Val) && isBST(root.Right, root.Val, max)
}

/*
临界值思路，
   root
   /  \
   l   r
  /\   / \
l1 l2  r1 r2
r 的临界值为[root,1<<63-1]
r1的临界值为[root, r]
*/
func isValidBST2(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	min, max := -1<<63, 1<<63-1
	return isBST(root, min, max)
}

func main() {
	root := &tree.TreeNode{Val: 0}
	//root := &tree.TreeNode{
	//	Val: 5,
	//	Left: &tree.TreeNode{
	//		Val: 2,
	//		Left: &tree.TreeNode{
	//			Val: 1,
	//		},
	//		Right: &tree.TreeNode{
	//			Val: 4,
	//		},
	//	},
	//	Right: &tree.TreeNode{
	//		Val: 8,
	//		Left: &tree.TreeNode{
	//			Val: 6,
	//		},
	//		Right: &tree.TreeNode{
	//			Val: 9,
	//		},
	//	},
	//}
	//fmt.Println(levelOrderTraversal1(root))
	fmt.Println(isValidBST2(root))
}
