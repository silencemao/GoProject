package main

import (
	"GoProject/leetcode/tree"
)

/*
给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。
修剪树不应该改变保留在树中的元素的相对结构（即，如果没有被移除，原有的父代子代关系都应当保留）。 可以证明，存在唯一的答案。
所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。
*/
func trimBST(root *tree.TreeNode, low int, high int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		right := trimBST(root.Right, low, high)
		return right
	}
	if root.Val > high {
		left := trimBST(root.Left, low, high)
		return left
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

func trimBST1(root *tree.TreeNode, low, high int) *tree.TreeNode {
	for root != nil {
		for root != nil && (root.Val < low || root.Val > high) {
			if root.Val < low {
				root = root.Right
			} else if root.Val > high {
				root = root.Left
			}
		}
	}

	cur := root
	for cur != nil {
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}

	cur = root
	for cur != nil {
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}

	return root
}

func main() {

}
