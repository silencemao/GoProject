package main

import (
	_struct "GoProject/leetcode/struct"
	"fmt"
	"strconv"
)

/*
是否为回文链表
*/
func isPalindrome234(head *_struct.ListNode) bool {
	str := ""
	for head != nil {
		str += strconv.Itoa(head.Val)
		head = head.Next
	}
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 { // 可能会超时
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

//快慢指针 反转前半部分链表

// 1->2->3->3->2->1->nil变成
// nil<-1<-2<-3 3->2->1-nil
func isPalindrome2341(head *_struct.ListNode) bool {
	slow, fast, pre := head, head, head
	var prepre *_struct.ListNode
	for fast != nil && fast.Next != nil {
		pre = slow
		slow, fast = slow.Next, fast.Next.Next

		pre.Next = prepre
		prepre = pre
	}
	if fast != nil {
		slow = slow.Next
	}

	for slow != nil && pre != nil {
		if slow.Val != pre.Val {
			return false
		}
		slow, pre = slow.Next, pre.Next
	}
	return true
}

func reverseL(head *_struct.ListNode) *_struct.ListNode {
	next := head
	var pre *_struct.ListNode
	for next != nil {
		tmp := next.Next
		next.Next = pre

		pre = next
		next = tmp
	}
	return pre
}

func main() {
	head := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 1}}}}
	fmt.Println(isPalindrome234(head))
}
