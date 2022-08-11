package main

import _struct "GoProject/leetcode/struct"

/*
寻找链表中间节点，快慢指针
*/
func middleNode(head *_struct.ListNode) *_struct.ListNode {
	s, f := head, head

	for f != nil && f.Next != nil {
		s, f = s.Next, f.Next.Next
	}

	return s
}

func main() {

}
