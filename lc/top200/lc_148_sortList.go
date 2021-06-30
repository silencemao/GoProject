package main

import _struct "GoProject/leetcode/struct"

/*
链表排序
归并排序，分治
*/
func sortList(head *_struct.ListNode) *_struct.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast, pre := head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil // 将链表的前后两部分断开
	return merge148(sortList(head), sortList(slow))
}

func merge148(l1, l2 *_struct.ListNode) *_struct.ListNode {
	dummy := &_struct.ListNode{Val: -1}
	tmp := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return dummy.Next
}

func main() {
	head := &_struct.ListNode{Val: 7, Next: &_struct.ListNode{
		Val: 6, Next: &_struct.ListNode{
			Val: 5, Next: &_struct.ListNode{
				Val: 4, Next: &_struct.ListNode{Val: 3}}}}}
	head = sortList(head)
	_struct.PrintList(head)
}
