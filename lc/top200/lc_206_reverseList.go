package main

import (
	_struct "myGoProject/leetcode/struct"
)

func reverseList(head *_struct.ListNode) *_struct.ListNode {
	var pre *_struct.ListNode
	next := head
	for next != nil {
		tmp := next.Next
		next.Next = pre
		pre = next
		next = tmp
	}
	return pre
}

func main() {
	l1 := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 4, Next: &_struct.ListNode{Val: 5}}}
	_struct.PrintList(l1)
	l1 = reverseList(l1)
	_struct.PrintList(l1)
}
