package main

import _struct "GoProject/leetcode/struct"

/*
链表插入排序
*/
func insertionSortList(head *_struct.ListNode) *_struct.ListNode {
	dummy := &_struct.ListNode{Val: -1, Next: head}
	lastSorted, cur := head, head.Next

	for cur != nil {
		if lastSorted.Val <= cur.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummy
			for prev.Next.Val <= cur.Val {
				prev = prev.Next
			}
			lastSorted.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}
		cur = lastSorted.Next
	}
	return dummy.Next
}

func main() {
	head := _struct.NewList([]int{3, 2, 1, 5, 4, 7})
	head = insertionSortList(head)
	_struct.PrintList(head)
}
