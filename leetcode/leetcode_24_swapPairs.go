package main

import "fmt"

type ListNode struct {
	value int
	next *ListNode
}

func generateListNode(array []int) *ListNode {
	head := new(ListNode)
	head.value = array[0]
	tmp := head

	for i := 1; i < len(array); i++ {
		node := new(ListNode)
		node.value = array[i]
		tmp.next = node
		tmp = tmp.next
	}
	return head
}

func printListNode(node *ListNode) {
	tmp := node
	for tmp.next != nil {
		fmt.Print(tmp.value, " ")
		tmp = tmp.next
	}
	fmt.Println(tmp.value)
}
/*
因为节点两两交换是涉及到3个节点的，pre、cur、next，所以我们需要先给整个链表添加一个"头"， 这样相对前两个节点就可以生成一个pre了
交换的过程就很简单了，先让pre指向next，再让cur指向next.next，最后让next指向cur就行。
交换完两个节点，再更新pre，cur，next的时候，pre要等于cur，而不是next，这个地方要注意。以及要判断节点是否为nil
*/
func swapPairs(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.value = 0
	dummy.next = head

	pre := dummy
	cur := dummy.next
	next := cur.next
	for cur != nil && next != nil {
		pre.next = next
		cur.next = next.next
		next.next = cur

		pre = cur
		cur = cur.next
		if cur != nil {
			next = cur.next
		}

	}
	return dummy.next
}

func main() {
	var array = []int{1, 2, 3, 4, 5, 6}
	head := generateListNode(array)
	printListNode(head)

	head = swapPairs(head)

	printListNode(head)

}