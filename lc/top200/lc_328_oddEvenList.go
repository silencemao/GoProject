package main

import (
	_struct "GoProject/leetcode/struct"
)

/*
奇偶链表
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
*/
func oddEvenList(head *_struct.ListNode) *_struct.ListNode {
	if head == nil {
		return head
	}
	odd := head
	evenHead := head.Next
	even := head.Next

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		even.Next = even.Next.Next

		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func main() {
	head := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{
		Val: 2, Next: &_struct.ListNode{
			Val: 3, Next: &_struct.ListNode{
				Val: 4, Next: &_struct.ListNode{
					Val: 5, Next: &_struct.ListNode{
						Val: 6, Next: &_struct.ListNode{
							Val: 7}}}}}}}

	head = oddEvenList(head)
	_struct.PrintList(head)
}
