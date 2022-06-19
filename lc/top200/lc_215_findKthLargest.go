package main

import (
	"container/heap"
	"fmt"
)

func partition(nums []int, left, right int) int {
	tmp := nums[left]
	for left < right {
		for left < right && nums[right] <= tmp {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] >= tmp {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = tmp
	return left
}

func helper(nums []int, left, right, k int) int {
	i := partition(nums, left, right)
	cnt := i - left + 1
	if cnt == k {
		return nums[i]
	} else if cnt > k {
		return helper(nums, left, i-1, k)
	} else {
		return helper(nums, i+1, right, k-cnt)
	}
}
func findKthLargest(nums []int, k int) int {
	res := helper(nums, 0, len(nums)-1, k)
	return res
}

type ns []int

func (n ns) Len() int            { return len(n) }
func (n ns) Less(i, j int) bool  { return n[i] < n[j] }
func (n ns) Swap(i, j int)       { n[i], n[j] = n[j], n[i] }
func (n *ns) Push(v interface{}) { *n = append(*n, v.(int)) }
func (n *ns) Pop() interface{}   { a := *n; v := a[len(a)-1]; *n = a[:len(a)-1]; return v }

func findKthLargest2(nums []int, k int) int {
	n := &ns{}
	for _, num := range nums {
		heap.Push(n, num)
		if n.Len() > k {
			heap.Pop(n)
		}
	}
	res := heap.Pop(n).(int)
	fmt.Println(res)
	return res
}
func main() {
	nums := []int{9, 8, 0, 1, 4, 3, 9, 6, 5, 3}
	k := 4
	//fmt.Println(findKthLargest(nums, k))
	//fmt.Println(findKthLargest1(nums, k))
	fmt.Println(findKthLargest2(nums, k))
}
