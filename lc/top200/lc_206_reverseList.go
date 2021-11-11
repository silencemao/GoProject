package main

import (
	_struct "GoProject/leetcode/struct"
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

func help206(pre, cur *_struct.ListNode) *_struct.ListNode {
	if cur == nil {
		return pre
	}
	tmp := cur.Next
	cur.Next = pre

	return help206(cur, tmp)
}

func reverseList1(head *_struct.ListNode) *_struct.ListNode {
	return help206(nil, head)
}

func reverseList2(head *_struct.ListNode) *_struct.ListNode {
	var pre *_struct.ListNode // 只声明类型 不初始化
	cur := head
	for cur != nil {
		tmp := cur.Next

		cur.Next = pre

		pre = cur
		cur = tmp
	}
	return pre
}

func main() {
	l1 := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 4, Next: &_struct.ListNode{Val: 5}}}
	_struct.PrintList(l1)
	l1 = reverseList(l1)
	_struct.PrintList(l1)
}
