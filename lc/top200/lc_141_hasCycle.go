package main

import _struct "GoProject/leetcode/struct"

// 链表是否有环，快慢指针
func hasCycle(head *_struct.ListNode) bool {
	l, f := head, head
	for f != nil && f.Next != nil {
		l = l.Next
		f = f.Next.Next
		if l == f {
			return true
		}
	}
	return false
}

func main() {

}
