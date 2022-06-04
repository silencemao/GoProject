package main

import (
	_struct "GoProject/leetcode/struct"
)

/*
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
删除链表中的重复元素
输入 1 1 2 3 3
返回 2
和第83题不同
*/
//删除以 head 作为开头的有序链表中，值出现重复的节点。
func deleteDuplicates82(head *_struct.ListNode) *_struct.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val != head.Next.Val {
		head.Next = deleteDuplicates82(head.Next) // 二者不相等，则以head.next作为新的开头判断
	} else { //相等，则找到第一个和head不相等的节点，作为开头判断
		move := head.Next
		for move != nil && head.Val == move.Val {
			move = move.Next
		}
		return deleteDuplicates82(move) // move后面
	}
	return head
}

func deleteDuplicates82II(head *_struct.ListNode) *_struct.ListNode {
	dummy := &_struct.ListNode{Val: -1, Next: head}
	pre, cur := dummy, head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if pre.Next == cur { // pre和cur之间没有间隔
			pre = pre.Next
		} else {
			pre.Next = cur.Next // 此时pre不能动，因为不确定pre.Next 和pre.Next.Next的关系
		}
		cur = cur.Next
	}
	return dummy.Next
}

func main() {
	head := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 3, Next: &_struct.ListNode{Val: 4}}}}}
	_struct.PrintList(head)
	head = deleteDuplicates82(head)
	_struct.PrintList(head)
}
