package main

import "fmt"

/*
旋转数据求最小值
已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
你必须设计一个时间复杂度为O(log n) 的算法解决此问题。
*/
func findMin(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] > nums[r] { // 最小值在右边
			l = mid + 1
		} else { // 最小值在左边
			r = r - 1
		}
	}
	return nums[l] // 最后一步 l=r=mid，然后nums[mid] = nums[l] --> r-=1，l所指的位置为最小值
}

func main() {
	nums := []int{11, 13, 15, 17}
	fmt.Println(findMin(nums))
}
