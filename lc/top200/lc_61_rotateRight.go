package main

import _struct "GoProject/leetcode/struct"

/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动k个位置。

输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

1、计算链表长度cnt，对k取余数
2、判断边界条件，比如链表为空或k=0
3、快慢指针，fast先走k步，然后slow 和 fast一起走，走到fast.next=nil的情况下
4、这时slow指在倒数第k+1个节点位置
5、把slow后面的部分放到头部即可
*/
func rotateRight(head *_struct.ListNode, k int) *_struct.ListNode {
	if head == nil {
		return head
	}
	cnt := 0
	tmp := head
	for tmp != nil {
		cnt += 1
		tmp = tmp.Next
	}
	k = k % cnt
	if k == 0 {
		return head
	}
	slow, fast := head, head
	for k > 0 && fast != nil {
		fast = fast.Next
		k -= 1
	}
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	tmp = slow.Next
	slow.Next = nil
	fast.Next = head
	return tmp
}

func main() {
	head := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 3, Next: &_struct.ListNode{Val: 4, Next: &_struct.ListNode{Val: 5, Next: nil}}}}}
	_struct.PrintList(head)

	tmp := rotateRight(head, 2)
	_struct.PrintList(tmp)

}
