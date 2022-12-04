package main

import (
	"container/heap"
	"fmt"
)

/*
查找和最小的K对数字
给定两个以 升序排列 的整数数组 nums1 和 nums2,以及一个整数 k。

定义一对值(u,v)，其中第一个元素来自nums1，第二个元素来自 nums2。

请找到和最小的 k个数对(u1,v1), (u2,v2) ... (uk,vk)。

*/
type pr struct{ i, j int }
type hp373 struct {
	ind          []pr
	nums1, nums2 []int
}

func (h hp373) Len() int { return len(h.ind) }
func (h hp373) Less(i, j int) bool {
	a, b := h.ind[i], h.ind[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h hp373) Swap(i, j int) { h.ind[i], h.ind[j] = h.ind[j], h.ind[i] }

func (h *hp373) Push(v interface{}) { h.ind = append(h.ind, v.(pr)) }
func (h *hp373) Pop() interface{}   { a := h.ind; v := a[len(a)-1]; h.ind = a[:len(a)-1]; return v }

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var res [][]int
	m, n := len(nums1), len(nums2)
	h := &hp373{nums2: nums2, nums1: nums1}
	for i := 0; i < k && i < m; i++ { // 存储nums1[i]和nums2[0]的组合
		heap.Push(h, pr{i: i, j: 0})
		//h.ind = append(h.ind, pr{i: i, j: 0})
	}

	for h.Len() > 0 && len(res) < k {
		p := heap.Pop(h).(pr)
		x, y := p.i, p.j
		res = append(res, []int{nums1[x], nums2[y]})
		if y+1 < n { // nums2向右扩展 nums2[y+1]
			heap.Push(h, pr{x, y + 1})
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}
