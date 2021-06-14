package main

import (
	"fmt"
	"sort"
)

/*
摆动排序，给定一个数组，将数组中的元素按照nums[0] < nums[1] > nums[2] < nums[3]的顺序进行排序
1、首先将数组从小到大排序 排序之后数组后半部分的大小都大于前半部分的数字大小
2、将后半部分的数字与前半部分的数字进行交叉摆放。
3、先填充奇数部分，这一部分的数字都是大数，我们先将原数组的后半部分的数填充到奇数位置。
4、再将数组前半部分的数字填充到偶数部分。
再填充的过程中，可以利用 索引k 来 逆序 取数组的数字，放置到tmp数组的奇数/偶数部分

同第280题
*/
func wiggleSort(nums []int) {
	sort.Ints(nums)
	tmp := make([]int, len(nums))
	k := len(nums)
	if k < 2 {
		return
	}
	for i := 1; i < len(nums); i += 2 {
		k--
		tmp[i] = nums[k]
	}
	for i := 0; i < len(nums); i += 2 {
		k--
		tmp[i] = nums[k]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = tmp[i]
	}
}

/*
奇数/偶数部分 递增
但这种写法 会存在问题，即数组中数字存在重复时不对
[4, 5, 5, 6] -> [4, 5, 6, 6]
*/
func wiggleSort1(nums []int) {
	sort.Ints(nums)
	tmp := make([]int, len(nums))
	k := len(nums) / 2
	if len(nums)&1 == 1 { // nums大小为奇数
		k += 1
	}
	for i := 1; i < len(nums); i += 2 {
		tmp[i] = nums[k]
		k++
	}
	k = 0
	for i := 0; i < len(nums); i += 2 {
		tmp[i] = nums[k]
		k++
	}
	fmt.Println(tmp)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	wiggleSort(nums)
	fmt.Println(nums)
}
