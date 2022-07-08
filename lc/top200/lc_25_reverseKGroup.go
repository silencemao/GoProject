package main

import (
	_struct "GoProject/leetcode/struct"
)

func h25(head, end *_struct.ListNode) {
	var dummy *_struct.ListNode
	pre, cur := dummy, head
	for cur != end {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}
}

/*
k个一组翻转链表
*/
func reverseKGroup(head *_struct.ListNode, k int) *_struct.ListNode {
	dummy := &_struct.ListNode{Val: -1, Next: head}
	pre, start, end := dummy, head, dummy
	for start != nil {
		i := 0
		for ; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil || i != k {
			break
		}
		next := end.Next
		h25(start, end.Next)
		//pre.Next = end
		//start.Next = next
		pre.Next, start.Next = end, next
		pre, start, end = start, next, start

	}
	return dummy.Next
}

func main() {
	head := _struct.NewList([]int{1, 2, 3, 4, 5, 6})
	k := 2
	_struct.PrintList(head)
	head = reverseKGroup(head, k)
	_struct.PrintList(head)
}
