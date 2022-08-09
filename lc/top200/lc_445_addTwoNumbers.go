package main

import _struct "GoProject/leetcode/struct"

/*
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

链表相加
*/
func addTwoNumbers445(l1 *_struct.ListNode, l2 *_struct.ListNode) *_struct.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var s1, s2 []*_struct.ListNode

	head := &_struct.ListNode{Val: 0}
	tail := head
	for l1 != nil {
		s1 = append(s1, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2)
		l2 = l2.Next
	}
	carry := 0
	var res []int
	for len(s1) > 0 && len(s2) > 0 {
		tmp := carry + s1[len(s1)-1].Val + s2[len(s2)-1].Val
		carry, tmp = tmp/10, tmp%10
		res = append(res, tmp)
		s1, s2 = s1[:len(s1)-1], s2[:len(s2)-1]
	}
	for len(s1) > 0 {
		tmp := carry + s1[len(s1)-1].Val
		carry, tmp = tmp/10, tmp%10
		res = append(res, tmp)
		s1 = s1[:len(s1)-1]
	}

	for len(s2) > 0 {
		tmp := carry + s2[len(s2)-1].Val
		carry, tmp = tmp/10, tmp%10
		res = append(res, tmp)
		s2 = s2[:len(s2)-1]
	}
	if carry > 0 {
		res = append(res, carry)
	}

	for i := len(res) - 1; i >= 0; i-- {
		head.Next = &_struct.ListNode{Val: res[i], Next: nil}
		head = head.Next
	}
	return tail.Next
}

func addTwoNumbers4452(l1 *_struct.ListNode, l2 *_struct.ListNode) *_struct.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var s1, s2 []*_struct.ListNode

	head := &_struct.ListNode{Val: 0}
	tail := head
	for l1 != nil {
		s1 = append(s1, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2)
		l2 = l2.Next
	}
	carry := 0
	var res []int
	for len(s1) > 0 || len(s2) > 0 || carry > 0 {
		a, b := 0, 0
		if len(s1) > 0 {
			a = s1[len(s1)-1].Val
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			b = s2[len(s2)-1].Val
			s2 = s2[:len(s2)-1]
		}
		tmp := carry + a + b
		carry, tmp = tmp/10, tmp%10
		res = append(res, tmp)
	}

	for i := len(res) - 1; i >= 0; i-- {
		head.Next = &_struct.ListNode{Val: res[i], Next: nil}
		head = head.Next
	}
	return tail.Next
}

func main() {
	l1 := _struct.NewList([]int{7, 2, 4, 3})
	l2 := _struct.NewList([]int{5, 6, 4})
	//head := addTwoNumbers445(l1, l2)
	head := addTwoNumbers4452(l1, l2)
	_struct.PrintList(head)
}
