package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

func help230(root *tree.TreeNode, res *[]int) {
	if root == nil {
		return
	}
	help230(root.Left, res)
	*res = append(*res, root.Val)
	help230(root.Right, res)
}

/*
二叉搜索树中第k小的元素
利用二叉搜索数的特性，中序遍历是单调递增的。中序遍历得到二叉搜索树的各节点的值，返回第k-1个即可
*/
func kthSmallest230(root *tree.TreeNode, k int) int {
	var res []int
	help230(root, &res)
	fmt.Println(res)
	return res[k-1]
}

/*
迭代中序遍历，校验res的长度，提前结束判断
*/
func kthSmallest2302(root *tree.TreeNode, k int) int {
	var res []int
	var stack []*tree.TreeNode
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			front := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, front.Val)
			root = front.Right
		}
		if len(res) == k {
			break
		}
	}
	return res[k-1]
}

func main() {
	root := &tree.TreeNode{Val: 5, Left: &tree.TreeNode{Val: 3, Left: &tree.TreeNode{Val: 2}, Right: &tree.TreeNode{Val: 4}}, Right: &tree.TreeNode{Val: 6, Right: &tree.TreeNode{Val: 7}}}
	fmt.Println(kthSmallest2302(root, 4))
}
