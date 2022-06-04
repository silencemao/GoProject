package _struct

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(arr []int) *ListNode {
	dummy := &ListNode{Val: -1}
	tmp := dummy
	for _, num := range arr {
		tmp.Next = &ListNode{Val: num}
		tmp = tmp.Next
	}
	return dummy.Next
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	dummy := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			head.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}
	return dummy.Next
}
func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print(" ")
		head = head.Next
	}
	fmt.Println()
}
