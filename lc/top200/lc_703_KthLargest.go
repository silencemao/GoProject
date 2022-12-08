package main

import "container/heap"

/*
数据流中的第K大元素

采用小根堆解决该问题
小根堆的大小为K，不断向里面添加元素，一旦数据个数>K，则弹出堆顶的元素。始终保持堆的第一个元素为最小元素，也就是第K大的元素
*/
type KthLargest struct {
	h *hp1046
	k int
}

func Constructor703(k int, nums []int) KthLargest {
	kl := KthLargest{h: &hp1046{}, k: k}

	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)

	if this.h.Len() > this.k {
		heap.Pop(this.h)
	}
	res := (*this.h)[0]

	return res
}

type hp703 []int

func (h hp703) Len() int           { return len(h) }
func (h hp703) Less(i, j int) bool { return h[i] < h[j] }
func (h hp703) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *hp703) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *hp703) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {

}
