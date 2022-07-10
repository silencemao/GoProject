package main

import (
	_struct "GoProject/leetcode/struct"
	"GoProject/leetcode/tree"
)

func help109(head *_struct.ListNode, end *_struct.ListNode) *tree.TreeNode {
	if head == nil || head == end {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast != end && fast.Next != end {
		slow, fast = slow.Next, fast.Next.Next
	}
	root := &tree.TreeNode{Val: slow.Val}
	root.Left = help109(head, slow)
	root.Right = help109(slow.Next, end)
	return root
}

/*
排序链表转为二叉搜索树
*/
func sortedListToBST(head *_struct.ListNode) *tree.TreeNode {
	root := help109(head, nil)
	return root
}

func main() {
	head := _struct.NewList([]int{-10, -3, 0, 5, 9})
	root := sortedListToBST(head)
	tree.LevelOrder(root)
}
