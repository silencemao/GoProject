package main

import _struct "GoProject/leetcode/struct"

/*
求两个链表的交点

     A -> B -> C \
                   D -> E -> F
M -> N -> O- > p /

从上图中可以看出D是两个链表的交点，交点前的部分长度是不一样的，如何弥补前面的差距就能找到交点
1、假如我们只是孤立的遍历两个链表，是无法直接找到交点的。
2、怎样在相同的时间内，让两个指针l、r走的步数一样。当l走到尽头时，它接着去headA处，r走到尽头时，它去headB处，这样就能弥补
较短的链表走的步数少，继续向前走 二者就会相交。
假设 headA 长度为a
     headB 长度为b 公共长度为c
a+(b-c) = b + (a-c)
b-c、a-c分别代表交点前的长度，上面的等式表示分别走完各自的链表长度，再去走另一个链表交点前的长度就会相遇
*/
func getIntersectionNode(headA, headB *_struct.ListNode) *_struct.ListNode {
	l, r := headA, headB
	for l != r {
		if l == nil {
			l = headB
		} else {
			l = l.Next
		}
		if r == nil {
			r = headA
		} else {
			r = r.Next
		}
	}
	return l
}

func main() {

}
