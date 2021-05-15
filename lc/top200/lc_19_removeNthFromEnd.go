package main

/*
删除链表的倒数第n个节点 并返回头节点
*/
func removeNthFromEnd(head *Node, n int) *Node {
	dummy := &Node{val: 0, next: head}
	l, r := dummy, dummy
	for n > 0 {
		r = r.next
		n--
	}
	for r.next != nil {
		l, r = l.next, r.next
	}
	if l.next != nil {
		l.next = l.next.next
	}
	return dummy.next
}

func main() {
}
