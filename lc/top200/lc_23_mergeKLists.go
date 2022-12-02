package main

import (
	_struct "GoProject/leetcode/struct"
	"container/heap"
	"fmt"
	"sort"
)

// 合并k个升序链表
func mergeKLists(lists []*_struct.ListNode) *_struct.ListNode {
	var tmp *_struct.ListNode
	for i := 0; i < len(lists); i++ {
		tmp = _struct.MergeTwoLists(tmp, lists[i])
	}
	return tmp
}

func merge(lists []*_struct.ListNode, l, r int) *_struct.ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return _struct.MergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeKLists1(lists []*_struct.ListNode) *_struct.ListNode {
	return merge(lists, 0, len(lists)-1)
}

func mergeKLists2(lists []*_struct.ListNode) *_struct.ListNode {
	var nodes []*_struct.ListNode
	for _, l := range lists {
		if l != nil {
			nodes = append(nodes, l)
		}
	}
	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})

	head := &_struct.ListNode{Val: 0}
	dummy := head
	for len(nodes) > 0 {
		dummy.Next = &_struct.ListNode{Val: nodes[0].Val}
		if nodes[0].Next != nil {
			nodes = append(nodes, nodes[0].Next)
		}
		nodes = nodes[1:]
		sort.SliceStable(nodes, func(i, j int) bool {
			return nodes[i].Val < nodes[j].Val
		})
		dummy = dummy.Next
	}
	return head.Next
}

func mergeKLists3(lists []*_struct.ListNode) *_struct.ListNode {
	pq := make(_struct.PriorityQueue, 0)
	for i, l := range lists {
		if l != nil {
			item := &_struct.Item{Value: l, Priority: l.Val, Index: i}
			pq = append(pq, item)
		}
	}
	heap.Init(&pq)
	head := &_struct.ListNode{Val: 0}
	dummy := head
	for pq.Len() > 0 {
		tmp := heap.Pop(&pq)
		dummy.Next = tmp.(*_struct.Item).Value
		if tmp.(*_struct.Item).Value.Next != nil {
			item := &_struct.Item{Value: tmp.(*_struct.Item).Value.Next, Priority: tmp.(*_struct.Item).Value.Next.Val}
			pq.Push(item)
		}
		dummy = dummy.Next
	}
	return head.Next
}

type ls []*_struct.ListNode

func (l ls) Len() int           { return len(l) }
func (l ls) Less(i, j int) bool { return l[i].Val < l[j].Val } // 小根堆
func (l ls) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func (l *ls) Push(v interface{}) { *l = append(*l, v.(*_struct.ListNode)) }
func (l *ls) Pop() interface{}   { a := *l; v := a[len(a)-1]; *l = a[:len(a)-1]; return v } // 弹出根节点

func f23(lists []*_struct.ListNode) *_struct.ListNode {
	head := &_struct.ListNode{Val: 0, Next: nil}
	h := &ls{}

	for _, l := range lists {
		if l == nil {
			continue
		}
		heap.Push(h, l)
	}
	dummy := head
	for h.Len() > 0 {
		tmp := heap.Pop(h).(*_struct.ListNode)
		dummy.Next = tmp
		dummy = dummy.Next
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}
	}
	return head.Next
}

func main() {
	l1 := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 4, Next: &_struct.ListNode{Val: 5}}}
	l2 := &_struct.ListNode{Val: 1, Next: &_struct.ListNode{Val: 3, Next: &_struct.ListNode{Val: 4}}}
	l3 := &_struct.ListNode{Val: 2, Next: &_struct.ListNode{Val: 6}}

	lists := []*_struct.ListNode{l1, l2, l3}
	head := mergeKLists(lists)
	_struct.PrintList(head)

	head1 := mergeKLists1(lists)
	_struct.PrintList(head1)
	//
	head2 := mergeKLists2(lists)
	_struct.PrintList(head2)

	head3 := mergeKLists3(lists)
	_struct.PrintList(head3)

	fmt.Println()
	head4 := f23(lists)
	_struct.PrintList(head4)
}
