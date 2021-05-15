package main

import (
	"fmt"
	_struct "GoProject/leetcode/struct"
)

func main() {
	l1 := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 4}}}
	l2 := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 3, Next: &_struct.ListNode{Val: 4, Next: &_struct.ListNode{Val: 8}}}}
	head := _struct.MergeTwoLists(l1, l2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
