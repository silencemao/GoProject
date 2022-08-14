package main

import _struct "GoProject/leetcode/struct"

/*
链表排序
归并排序，分治

自顶向下拆，全部拆完再合并，递归调用，空间复杂度高
*/
func sortList(head *_struct.ListNode) *_struct.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast, pre := head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil // 将链表的前后两部分断开
	return merge148(sortList(head), sortList(slow))
}

func merge148(l1, l2 *_struct.ListNode) *_struct.ListNode {
	dummy := &_struct.ListNode{Val: -1}
	tmp := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return dummy.Next
}

/*
自低向上，先拆长度为1的子段，合并，再拆长度为2的子段，合并....
1、拆长度为subL的两个字段
2、合并 两个字段
3、更新prev, cur，继续拆
*/
func sortList1(head *_struct.ListNode) *_struct.ListNode {
	l, tmp := 0, head
	for tmp != nil {
		l, tmp = l+1, tmp.Next
	}
	dummy := &_struct.ListNode{Val: -1, Next: head}
	for subL := 1; subL < l; subL = subL << 1 {
		prev, cur := dummy, dummy.Next
		for cur != nil {
			//取第一段
			head1 := cur
			for i := 1; i < subL && cur.Next != nil; i++ {
				cur = cur.Next
			}
			//取第二段
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subL && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// 合并
			var next *_struct.ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			prev.Next = merge148(head1, head2)
			// 更新
			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummy.Next
}

func main() {
	head := &_struct.ListNode{Val: 7, Next: &_struct.ListNode{
		Val: 6, Next: &_struct.ListNode{
			Val: 5, Next: &_struct.ListNode{
				Val: 4, Next: &_struct.ListNode{Val: 3}}}}}
	//head = sortList(head)
	head = sortList1(head)
	_struct.PrintList(head)
}
