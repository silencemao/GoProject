package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func topKFrequent692(words []string, k int) []string {
	set := make(map[string]int, 0)
	for _, w := range words {
		set[w] += 1
	}
	var uniqueWords []string
	for w := range set {
		uniqueWords = append(uniqueWords, w)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		w1, w2 := uniqueWords[i], uniqueWords[j]
		return set[w1] > set[w2] || (set[w1] == set[w2] && w1 < w2)
	})
	return uniqueWords[:k]
}

type pair struct {
	word string
	c    int
}

type hp []pair // 其实就是一个二叉树，hp[0]是根节点，后面依次是左右子树

func (h hp) Less(i, j int) bool {
	return h[i].c < h[j].c || (h[i].c == h[j].c && h[i].word > h[j].word)
}

func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h hp) Len() int { return len(h) }

func (h *hp) Pop() interface{} { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }

func topKFrequentII(words []string, k int) []string {
	set := make(map[string]int, 0)
	for _, w := range words {
		set[w] += 1
	}
	h := &hp{}
	//heap.Init(h)
	for w, c := range set {
		heap.Push(h, pair{word: w, c: c})
		for _, w := range *h {
			fmt.Print(w.word)
			fmt.Print(" ")
		}
		fmt.Println()
		if h.Len() > k {
			heap.Pop(h)
			//h.Pop() 这种写法就不对了，因为heap.Pop会先将堆顶和最后一个位置元素交换，再将堆顶删除，其实此处的堆顶就是二叉树的根节点
		}
	}
	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(pair).word
	}
	return res
}

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	fmt.Println(topKFrequent692(words, k))
	fmt.Println(topKFrequentII(words, k))
}
