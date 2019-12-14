package main

/*
去除链表中倒数第n个节点 n从1开始计数
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head==nil || n==0 {
		return head
	}

	count := 0
	tmp := head
	for tmp != nil {
		count += 1
		tmp = tmp.Next
	}

	if n > count {
		return head
	}
	if n==count {
		return head.Next
	}

	tmp = head
	num := count - n - 1
	for num > 0 && tmp != nil {
		num -= 1
		tmp = tmp.Next
	}
	if tmp.Next != nil {
		tmp.Next = tmp.Next.Next
	}
	return head
}
// 快慢指针
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if n <=0 {
		return head
	}
	fast := head
	low := head
	var i int
	for i = 0; i < n; i++ {
		if fast==nil {
			break
		}
		fast = fast.Next
	}
	// n > list length
	if i < n {
		return head
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		low = low.Next
		fast = fast.Next
	}

	low.Next = low.Next.Next
	return head
}

//func main() {
//	var array = []int{1, 2, 3, 4, 5}
//	head := generateListNode(array)
//	printListNode(head)
//	n := 5
//	head = removeNthFromEnd1(head, n)
//	printListNode(head)
//}
