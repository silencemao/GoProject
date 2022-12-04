package main

import (
	"container/heap"
	"fmt"
)

/*
给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。

1、利用一个队列记录窗口内的最大值的索引
2、当队列中的数字个数大于k时，去除队首的元素
3、当有新的元素加入到队列中时，从队尾开始，逐个比较队列中的索引对应的num值 与 nums[i]的关系，
若小于nums[i]，删除队列中的该索引，遇到第一个不小于nums[i]的索引，即退出比较。 队列中索引对应的元素值是【递减的】
4、将索引i加入到队列中
5、若i>=k-1 表示 窗口已经覆盖了k个数，取队首加入到res中

*/
func maxSlidingWindow1(nums []int, k int) []int {
	var res, queue []int

	for i := 0; i < len(nums); i++ {
		if len(queue) > 0 && queue[0] == i-k { // queue中只保留 [i-k+1, ...i-1, i] k个数
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[i] { // 若新加入数字，只保留 >=nums[i] 的数
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i) // i 加入queue
		if i >= k-1 {
			res = append(res, nums[queue[0]]) // 取队首 队首一定 >= nums[i]
		}
	}

	return res
}

func maxSlidingWindow2(nums []int, k int) []int {
	var res, queue []int
	for i := 0; i < len(nums); i++ {
		if len(queue) > 0 && i-queue[0] >= k { // 队列里面始终位置i-k+1到i之间元素的索引
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[i] { // 删除i-k+1到i之间比nums[i]小的数字的索引
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}

/*
滑动窗口内最小的值
*/
func minSlidingWindow2(nums []int, k int) []int {
	var res, queue []int

	for i := 0; i < len(nums); i++ {
		if len(queue) > 0 && i-queue[0] >= k {
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[queue[len(queue)-1]] >= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i+1 >= k {
			res = append(res, nums[queue[0]])
		}
	}

	return res
}

func minSlidingWindow3(nums []int, k int) []int {
	var res []int
	for i := 0; i <= len(nums)-k; i++ {
		tmp := 1<<31 - 1
		for j := i; j < i+k; j++ {
			if nums[j] < tmp {
				tmp = nums[j]
			}
		}
		res = append(res, tmp)
	}
	return res
}

/*
优先队列
1、构造数据结构 lc239 有两个字段 ind 和 num
2、首先将元素像优先队列中放置，直到队列中元素个数等于k个
3、如果队列头部元素的索引<i-k，则弹出，因为它不在元素i合适的滑动窗口内
4、取出队列头部元素，注意此处不是弹出
*/
type lc239 struct {
	ind int
	num int
}

type ms []*lc239

func (m ms) Len() int           { return len(m) }
func (m ms) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ms) Less(i, j int) bool { return m[i].num > m[j].num }

func (m *ms) Push(v interface{}) { *m = append(*m, v.(*lc239)) }
func (m *ms) Pop() interface{}   { a := *m; v := a[len(a)-1]; *m = a[:len(a)-1]; return v }

func f239(nums []int, k int) []int {
	var res []int
	h := &ms{}
	for i := 0; i < len(nums); i++ {
		heap.Push(h, &lc239{ind: i, num: nums[i]})
		if h.Len() >= k {
			for (*h)[0].ind <= i-k {
				heap.Pop(h)
			}

			top := (*h)[0].num // 不是弹出，是取元素值
			res = append(res, top)
		}
	}
	return res
}

func main() {
	nums := []int{7, 4, 5, 1, 2, 3, 7, 8}
	k := 2
	fmt.Println(maxSlidingWindow1(nums, k))
	//fmt.Println(maxSlidingWindow2(nums, k))
	//fmt.Println(minSlidingWindow2(nums, k))
	//fmt.Println(minSlidingWindow3(nums, k))
	fmt.Println(f239(nums, k))
}
