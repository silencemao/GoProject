package main

import (
	_struct "GoProject/leetcode/struct"
	"fmt"
)

/*
分割链表
给你一个头结点为 head 的单链表和一个整数 k ，请你设计一个算法将链表分隔为 k 个连续的部分。
每部分的长度应该尽可能的相等：任意两部分的长度差距不能超过 1 。这可能会导致有些部分为 null 。
这 k 个部分应该按照在链表中出现的顺序排列，并且排在前面的部分的长度应该大于或等于排在后面的长度。

返回一个由上述 k 部分组成的数组。

链表长度为n，拆成k个部分，每部分长度为n/k，剩余n%k部分分别添加到前n%k个部分中，每部分里面多1个节点
*/
func splitListToParts(head *_struct.ListNode, k int) []*_struct.ListNode {
	var res, tmp []*_struct.ListNode
	for head != nil {
		tmp = append(tmp, head)
		head = head.Next
	}
	if len(tmp) < k {
		var a *_struct.ListNode
		for _, t := range tmp {
			t.Next = nil
		}
		for k > len(tmp) {
			tmp = append(tmp, a)
		}
		return tmp
	} else {
		base := len(tmp) / k
		add := len(tmp) % k
		i, ind := 0, 0
		for ; i < add; i++ {
			res = append(res, tmp[ind])
			tmp[ind+base].Next = nil
			ind += base + 1
		}
		for ; i < k; i++ {
			res = append(res, tmp[ind])
			tmp[ind+base-1].Next = nil
			ind += base
		}
	}

	return res
}

func main() {
	head := _struct.NewList([]int{1, 2, 3})
	k := 5
	res := splitListToParts(head, k)
	for _, r := range res {
		fmt.Println(r.Val)
	}
}
