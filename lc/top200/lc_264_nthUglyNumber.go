package main

import (
	"GoProject/leetcode/util"
	"container/heap"
	"fmt"
)

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p1, p2, p3 := 1, 1, 1

	for i := 2; i <= n; i++ {
		num2, num3, num5 := dp[p1]*2, dp[p2]*3, dp[p3]*5
		tmp := util.MinInt(util.MinInt(num2, num3), num5)
		dp[i] = tmp
		if tmp == num2 {
			p1 += 1
		}
		if tmp == num3 {
			p2 += 1
		}
		if tmp == num5 {
			p3 += 1
		}
	}
	return dp[n]
}

/*
优先队列
注意会存在重复的元素
*/
type lc264 []int

func (h lc264) Len() int            { return len(h) }
func (h lc264) Less(i, j int) bool  { return h[i] < h[j] }
func (h lc264) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *lc264) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *lc264) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func nthUglyNumber2(n int) int {
	h := &lc264{}
	res := 0
	heap.Push(h, 1)
	for n > 0 {
		top := heap.Pop(h).(int)
		heap.Push(h, top*2)
		heap.Push(h, top*3)
		heap.Push(h, top*5)
		res = top
		for h.Len() > 0 && (*h)[0] == top { // 去重
			heap.Pop(h)
		}
		n -= 1
	}
	return res
}

func main() {
	n := 15
	fmt.Println(nthUglyNumber(n))
	fmt.Println(nthUglyNumber2(n))
}
