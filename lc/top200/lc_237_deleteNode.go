package main

import _struct "GoProject/leetcode/struct"

/*
删除链表中的节点，只给出要删除的节点，不给出head
*/
func deleteNode(node *_struct.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	head := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 3, Next: &_struct.ListNode{Val: 1}}}}
	deleteNode(head.Next)
}
