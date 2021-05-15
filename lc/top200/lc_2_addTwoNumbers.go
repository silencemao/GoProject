package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func addTwoNumbers(l, r *Node) *Node {
	carry := 0
	head := new(Node)
	tail := head
	a, b, sum := 0, 0, 0
	for l != nil && r != nil {
		if l != nil {
			a = l.val
			l = l.next
		}
		if r != nil {
			b = r.val
			r = r.next
		}
		sum = a + b + carry
		carry = sum / 10
		head.next = &Node{val: sum % 10}
		head = head.next
	}
	for l != nil {
		a = l.val
		sum = a + carry
		head.next = &Node{val: sum % 10}
		carry = sum % 10
		l = l.next
	}
	for r != nil {
		a = r.val
		sum = a + carry
		head.next = &Node{val: sum % 10}
		carry = sum / 10
		r = r.next
	}
	if carry != 0 {
		head.next = &Node{val: carry}
	}
	return tail.next
}

func main() {
	l := &Node{val: 6, next: &Node{val: 3, next: &Node{val: 7}}}
	r := &Node{val: 6, next: &Node{val: 9, next: &Node{val: 8, next: &Node{val: 5}}}}

	res := addTwoNumbers(l, r)
	for res != nil {
		fmt.Print(res.val)
		res = res.next
	}
	fmt.Println()
	fmt.Println(1 + 2)
}
