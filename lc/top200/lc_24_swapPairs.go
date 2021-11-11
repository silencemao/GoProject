package main

import _struct "GoProject/leetcode/struct"
/*
两两交换链表中的节点
pre cur next
1、pre向后 找到一对的后者
2、cur向后 找到下一对的前者
3、next向前 找到当前对的前者

pre 一对节点之前的那个点 cur 一对节点的第一个点 next一对节点的第二个点
*/
func swapPairs(head *_struct.ListNode) *_struct.ListNode {
	dummy := &_struct.ListNode{Val: 0, Next: head}

	pre, cur := dummy, dummy.Next
	for cur != nil && cur.Next != nil {
		next := cur.Next

		pre.Next = cur.Next
		cur.Next = next.Next
		next.Next = cur

		pre = cur
		cur = cur.Next
	}

	return dummy.Next
}

func main() {
	head := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 3,Next: &_struct.ListNode{Val: 4}}}}
	head = swapPairs(head)

	_struct.PrintList(head)
}
