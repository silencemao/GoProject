package main

import "fmt"

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

func main() {
	nums := []int{7, 2, 4}
	k := 2
	fmt.Println(maxSlidingWindow1(nums, k))
}
