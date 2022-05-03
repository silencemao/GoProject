package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个 排序好 的数组arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。

整数 a 比整数 b 更接近 x 需要满足：
|a - x| < |b - x| 或者
|a - x| == |b - x| 且 a < b
*/
/*
数组两端的元素和x的差值大于等于数组中间元素和x的差值，所以我们可以考虑逐渐的删除数组两端的元素
删除几个？len(arr)-k个即可。
*/
func findClosestElements(arr []int, k int, x int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	for len(res) > k {
		l, r := 0, len(res)-1
		sub1, sub2 := util.Abs(res[l]-x), util.Abs(res[r]-x)
		if sub1 > sub2 {
			res = res[1:]
		} else {
			res = res[:r]
		}
	}
	return res
}

/*
上面方式的优化，不需要额外内存
*/
func findClosestElements1(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1
	delNum := len(arr) - k
	for delNum > 0 {
		sub1, sub2 := util.Abs(arr[l]-x), util.Abs(arr[r]-x)
		if sub1 > sub2 { // 删除左边
			l += 1
		} else { // 删除右边
			r -= 1
		}
		delNum -= 1
	}
	return arr[l : l+k]
}

/*
二分法，左边界的范围是0 ~ len(arr)-k，逐渐逼近左边界
*/

func findClosestElements2(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-k
	for l < r {
		mid := l + (r-l)>>1
		if x-arr[mid] > arr[mid+k]-x { //[mid,mid+k]之间左边界和x的距离更远，所以调整左边界
			l = mid + 1
		} else {
			r = mid
		}
	}
	return arr[l : l+k]
}

func main() {
	arr := []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}
	k, x := 3, 5
	fmt.Println(findClosestElements(arr, k, x))
	fmt.Println(findClosestElements1(arr, k, x))
	fmt.Println(findClosestElements2(arr, k, x))
}
