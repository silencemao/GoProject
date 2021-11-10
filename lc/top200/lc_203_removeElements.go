package main

import _struct "GoProject/leetcode/struct"
/*
链表 删除元素
*/
func removeElements(head *_struct.ListNode, val int) *_struct.ListNode {
	dummy := &_struct.ListNode{Val: 0, Next:head}

	cur := dummy
	for cur != nil && cur.Next != nil {
		next := cur.Next
		if next.Val == val {
			cur.Next = next.Next
		} else {
			cur = cur.Next // 这个要放在else里面，如果放在外面 对于 1-2-3-2-2-2这种情况就不对
		}
	}

	return dummy.Next
}

func main() {
	head := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 3, Next: &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 2}}}}}
	head = removeElements(head, 2)
	_struct.PrintList(head)

}
