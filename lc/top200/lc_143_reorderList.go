package main

import (
	_struct "GoProject/leetcode/struct"
)

/*
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

1-2-3-4
变成
1-4-2-3

方法一、将所有节点存储到res中，然后分别取节点重新连接
方法二、找中点 + 翻转后半段 + 合并前后两段
*/

func reorderList(head *_struct.ListNode) {
	var res []*_struct.ListNode
	tmp := head
	for tmp != nil {
		res = append(res, tmp)
		tmp = tmp.Next
	}
	var ppre, pre, cur *_struct.ListNode
	for i := 0; i < len(res)/2; i++ {
		pre, cur = res[i], res[len(res)-1-i]
		pre.Next, cur.Next = cur, nil

		if ppre != nil {
			ppre.Next = pre
		}
		ppre = cur
	}
	if len(res) > 1 && len(res)%2 == 1 {
		cur.Next = res[len(res)/2]
		res[len(res)/2].Next = nil
	}
}

func reorderList1(head *_struct.ListNode) {
	var res []*_struct.ListNode
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	i, j := 0, len(res)-1
	for i < j {
		res[i].Next = res[j]
		i += 1
		if i == j {
			break
		}
		res[j].Next = res[i]
		j -= 1
	}
	res[i].Next = nil
}

func findMid(head *_struct.ListNode) *_struct.ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
		s, f = s.Next, f.Next.Next
	}
	return s
}

func reverseLN(head *_struct.ListNode) *_struct.ListNode {
	var dummy *_struct.ListNode
	pre, cur := dummy, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}
	return pre
}

func mergeL(l1, l2 *_struct.ListNode) *_struct.ListNode {
	h1, h2 := l1, l2
	for h1 != nil && h2 != nil {
		h1N, h2N := h1.Next, h2.Next

		h1.Next = h2
		h1 = h1N

		h2.Next = h1N
		h2 = h2N
	}
	return l1
}

func reorderList2(head *_struct.ListNode) {
	mid := findMid(head)
	l1, l2 := head, mid.Next
	mid.Next = nil
	l2 = reverseLN(l2)

	mergeL(l1, l2)
}

func main() {
	head := _struct.NewList([]int{1, 2, 3, 4, 5, 6})
	_struct.PrintList(head)
	//reorderList(head)
	reorderList2(head)
	_struct.PrintList(head)
}
