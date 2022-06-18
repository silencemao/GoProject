package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
前K个高频元素
此题还可以用堆来做，学习中
*/
func topKFrequent(nums []int, k int) []int {
	if k > len(nums) {
		return []int{}
	}
	set := make(map[int]int, 0)
	for _, num := range nums {
		set[num] += 1
	}
	tmp := []int{}
	for num := range set {
		tmp = append(tmp, num)
	}
	sort.SliceStable(tmp, func(i, j int) bool {
		return set[tmp[i]] > set[tmp[j]]
	})

	return tmp[:k]
}

func topKFrequentI(nums []int, k int) []int {
	set := make(map[int]int, 0)
	for _, n := range nums {
		set[n] += 1
	}
	h := &hp1{}

	for w, c := range set {
		heap.Push(h, pair1{n: w, c: c})

		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(pair1).n
	}
	return res
}

type pair1 struct {
	n int
	c int
}

type hp1 []pair1

func (h hp1) Len() int           { return len(h) }
func (h hp1) Less(i, j int) bool { return h[i].c < h[j].c }
func (h hp1) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *hp1) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp1) Push(v interface{}) { *h = append(*h, v.(pair1)) }

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
