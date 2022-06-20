package main

import (
	"container/heap"
	"fmt"
)

/*
给定一个数组 points，其中points[i] = [xi, yi]表示 X-Y 平面上的一个点，并且是一个整数 k ，返回离原点 (0,0) 最近的 k 个点。

这里，平面上两点之间的距离是欧几里德距离（√(x1- x2)2+ (y1- y2)2）。

你可以按 任何顺序 返回答案。除了点坐标的顺序之外，答案 确保 是 唯一 的。

*/
type point struct {
	p []int
	d int
}

type ps []point

func (p ps) Len() int           { return len(p) }
func (p ps) Less(i, j int) bool { return p[i].d > p[j].d } // 不需要保留放在前面
func (p ps) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *ps) Push(v interface{}) { *p = append(*p, v.(point)) }
func (p *ps) Pop() interface{}   { a := *p; v := a[len(a)-1]; *p = a[:len(a)-1]; return v }

func dis(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}

func kClosest(points [][]int, k int) [][]int {
	res := make([][]int, k)

	h := &ps{}
	for p := range points {
		heap.Push(h, point{p: points[p], d: dis(points[p], []int{0, 0})})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(point).p
	}

	return res
}

func main() {
	points := [][]int{{1, 3}, {-2, 2}}
	k := 1
	fmt.Println(kClosest(points, k))
}
