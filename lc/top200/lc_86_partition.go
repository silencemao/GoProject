package main

import _struct "GoProject/leetcode/struct"

/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。

给定参考值x，将链表中小于x的节点放在大于等于x节点的前面，相对顺序保持不变
1、首先找到一个节点pre，pre.Next >= x，将pred后面所有小于x节点，逐个放在pre后面
2、cur=pre.Next，从这里开始找所有小于x的节点，注意这里要用双指针，cur,next，next为我们要判断的节点值是否小于x，cur为next的前驱
3、若next.val<x，则将next放到pre后面，更新pre，同时cur指向next.Next，注意此时cur不更新，因为cur.Next是新的节点，还没判断
4、若next.Val >= x，则更新cur

*/
func partition86(head *_struct.ListNode, x int) *_struct.ListNode {
	dummy := &_struct.ListNode{Val: -1, Next: head}
	pre, cur := dummy, head
	for pre != nil && pre.Next != nil {
		if pre.Next.Val < x {
			pre = pre.Next
		} else {
			break
		}
	}
	if pre == nil {
		return dummy.Next
	}
	cur = pre.Next
	for cur != nil {
		next := cur.Next
		if next != nil && next.Val < x {
			cur.Next = next.Next

			next.Next = pre.Next
			pre.Next = next
			pre = pre.Next
		} else {
			cur = cur.Next
		}

	}
	return dummy.Next
}

func main() {
	head := _struct.NewList([]int{2, 1})
	_struct.PrintList(head)

	head = partition86(head, 2)
	_struct.PrintList(head)
}
