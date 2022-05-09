package main

import (
	_struct "GoProject/leetcode/struct"
	"fmt"
)

/*
删除排序链表中的重复元素
*/
func deleteDuplicates(head *_struct.ListNode) *_struct.ListNode {
	dummy := &_struct.ListNode{Val: -101, Next: nil}
	dummy.Next = head
	tmp, next := dummy, dummy.Next
	for next != nil {
		if next.Val == tmp.Val {
			next = next.Next
		} else {
			tmp.Next = next
			tmp = tmp.Next
			next = next.Next
		}
	}
	tmp.Next = nil
	return dummy.Next

}
func main() {
	head := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 3, Next: &_struct.ListNode{Val: 3}}}}}
	fmt.Println(deleteDuplicates(head))
}
