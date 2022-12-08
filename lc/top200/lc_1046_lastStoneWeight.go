package main

import (
	"container/heap"
	"fmt"
)

/*
最后石头的大小
给定一堆石头，先拿两个最大的石头 x,y x<=y ，进行撞击，撞击之后 x=0 y=y-x。则最后剩余的石头重量是多少
*/
func lastStoneWeight(stones []int) int {
	h := &hp1046{}
	for _, num := range stones {
		heap.Push(h, num)
	}

	for h.Len() >= 2 {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		if x != y {
			heap.Push(h, abs(x-y))
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}
func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

type hp1046 []int

func (h hp1046) Len() int           { return len(h) }
func (h hp1046) Less(i, j int) bool { return h[i] > h[j] }
func (h hp1046) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *hp1046) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *hp1046) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
	stones := []int{1}
	fmt.Println(lastStoneWeight(stones))
}
