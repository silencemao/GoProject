package main

import _struct "GoProject/leetcode/struct"

func detectCycle(head *_struct.ListNode) *_struct.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

func detectCycle1(head *_struct.ListNode) *_struct.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

func main() {

}
