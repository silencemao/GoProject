package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
)

type lc506 struct {
	ind, score int
}
type hs506 []lc506

func (h hs506) Len() int            { return len(h) }
func (h hs506) Less(i, j int) bool  { return h[i].score > h[j].score }
func (h hs506) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hs506) Push(v interface{}) { *h = append(*h, v.(lc506)) }
func (h *hs506) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

/*
给定score，score[i]表示分数，分数越高，排名越靠前
前三名分别为 Gold Medal Sliver Medal Bronze Medal 4名之后的用字符串数字表示即可
堆排序 大根堆
*/
func findRelativeRanks(score []int) []string {
	n := len(score)
	answer := make([]string, n)
	h := &hs506{}
	for ind, s := range score {
		heap.Push(h, lc506{ind: ind, score: s})
	}
	cnt := 1
	for h.Len() > 0 {
		top := heap.Pop(h).(lc506)
		if cnt == 1 {
			answer[top.ind] = "Gold Medal"
		} else if cnt == 2 {
			answer[top.ind] = "Silver Medal"
		} else if cnt == 3 {
			answer[top.ind] = "Bronze Medal"
		} else {
			answer[top.ind] = strconv.Itoa(cnt)
		}
		cnt += 1
	}

	return answer
}

// 模拟
func findRelativeRanks1(score []int) []string {
	n := len(score)
	answer := make([]string, n)
	tmp := append([]int{}, score...)
	sort.Ints(tmp)
	set := map[int]int{}
	for i, num := range tmp {
		set[num] = n - i
	}

	for i, n := range score {
		cnt := set[n]
		if cnt == 1 {
			answer[i] = "Gold Medal"
		} else if cnt == 2 {
			answer[i] = "Silver Medal"
		} else if cnt == 3 {
			answer[i] = "Bronze Medal"
		} else {
			answer[i] = strconv.Itoa(cnt)
		}
	}
	return answer
}

func main() {
	score := []int{10, 3, 8, 9, 4}
	fmt.Println(findRelativeRanks(score))
	fmt.Println(findRelativeRanks1(score))
}
