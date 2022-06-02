package main

import (
	_struct "GoProject/leetcode/struct"
	"fmt"
)

/*
给你单链表的头指针 head 和两个整数left 和 right ，其中left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

1、找出left-1, left, right, right-1所对应的节点
2、对left和right中间的部分进行翻转
3、链接两头的部分
*/
func reverseBetween(head *_struct.ListNode, left int, right int) *_struct.ListNode {
	if left == right {
		return head
	}
	dummy := &_struct.ListNode{Val: 0, Next: head}
	cnt := 0
	var lP, l, r, rN *_struct.ListNode
	tmp := dummy
	for tmp != nil {
		if cnt == left-1 {
			lP = tmp
		}
		if cnt == left {
			l = tmp
		}
		if cnt == right {
			r = tmp
		}
		if cnt == right+1 {
			rN = tmp
		}
		tmp = tmp.Next
		cnt += 1
	}
	if lP == nil || l == nil || r == nil {
		return head
	}

	var pre *_struct.ListNode
	cur := l
	for cur != nil && cur != rN {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}
	lP.Next = pre
	l.Next = rN
	return dummy.Next
}

/*
不断的把被翻转的节点作为被翻转区间第一个节点
1->2->3->4->5->nil
left=2 right=4
第一步 1 3 2 4 5
第二步 1 4 3 2 5
*/
func reverseBetween1(head *_struct.ListNode, left int, right int) *_struct.ListNode {
	dummy := &_struct.ListNode{Val: -1, Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next // 被翻转的节点

		cur.Next = next.Next
		next.Next = pre.Next // 被翻转的节点的下一个 一定是 pre后面那个节点
		pre.Next = next

		//下面这个操作不对
		//pre.Next = next
		//cur.Next = next.Next
		//next.Next = pre.Next // 会导致pre 和 next互相指向对方

	}
	return dummy.Next
}

func main() {
	head := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 3, Next: &_struct.ListNode{Val: 4, Next: &_struct.ListNode{Val: 5}}}}}
	_struct.PrintList(head)
	fmt.Println(head.Val)
	head = reverseBetween1(head, 1, 5)
	_struct.PrintList(head)
}
