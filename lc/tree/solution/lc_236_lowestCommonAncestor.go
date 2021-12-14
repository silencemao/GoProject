package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

最近公共祖先的定义为：对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大,一个节点也可以是它自己的祖先
https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree

设节点 root 为节点 p,q 的某公共祖先，若其左子节点 root.left 和右子节点 root.right 都不是 p,q的公共祖先，则称 root 是 最近的公共祖先。

递归的遍历左右子树，返回值分别即为left， right。left表示左子树中可以作为p、q其中一个或两个祖先的一个节点，right同理。
若left right都不为空，表示left、right分别为p、q各自的祖先，则left、right的父节点是p、q的公共祖先
若left为nil， right不为nil，表示左子树中没有p、q中任意一个节点的祖先，right为它们的公共祖先，返回right即可。
若left不为nil，right为nil时，同上
若left为nil， right为nil，没有公共祖先
退出条件，某节点为nil 或 节点等于p 或 节点等q，表示当前节点可以为p或q其中一个的祖先

*/
func lowestCommonAncestor(root, p, q *tree.TreeNode) *tree.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil && right != nil {
		return right
	}
	if left != nil && right == nil {
		return left
	}
	return nil
}
func main() {
	root := tree.CreateTreeNode([]int{3, 5, 1, 6, 2, 0, 8, tree.NilNode, tree.NilNode, 7, 4})
	tree.LevelOrder(root)
	fmt.Println(lowestCommonAncestor(root, root.Left.Left, root.Left.Right.Right))
}
