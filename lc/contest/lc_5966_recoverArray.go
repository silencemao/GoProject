package main

import (
	"fmt"
	"sort"
)

/*
有一个数组arr(未知), 一个数字k(未知)，以及一个数组nums(已知)
nums的来源是由arr和k的变换得到的，变换方式如下
low[i] = arr[i]-k
high[i] = arr[i]+k
nums是由low和high合并之后的，所以nums的大小是arr的两倍

求出arr是什么

通过题意我们可以知道high[i]-low[i] = 2*k
以及low[0]=arr[0]-k（假设arr是从小到大排序)，low[0]肯定是nums中最小的数字了，但是arr[0]和k未知，只能遍历了，
也就是遍历nums中其它数和low[0]相减之后，差值是否为偶数 nums[i]-nums[0] ? 2*k
是偶数，表示nums[i]可以为high[0], 则可算出k为多少，并且也得到了arr[0]，arr中其它元素的计算方式也是这样
使用双指针遍历nums，存在nums[l]+nums[r] = 2*k，则表示又找到了一对low和high

几个注意的点：
1、要记录nums中已经使用过的数字，避免重复使用
2、nums[l]+nums[r] < 2*k时，r++
3、nums[l]+nums[r] > 2*k时，l++
对于下面第二种写法，若进入第2、3步要continue，因为还要判断l++/r++之后的nums中的数字是否能继续使用
*/
// [1 2 3 4 6 7 7 8 8 8 9 9 9 10 10 11]
func recoverArray(nums []int) []int {
	sort.Ints(nums)
	size := len(nums)
	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[0]
		if d == 0 || d%2 == 1 {
			continue
		}
		var res []int
		res = append(res, (nums[0]+nums[i])/2)
		used := make([]bool, len(nums))
		used[i] = true
		k := d / 2
		for l, r := 0, i+1; r < len(nums); r++ {
			for l++; l < size && used[l]; l++ {
			}
			for ; r < size && nums[r]-nums[l] < 2*k; r++ {
			}
			if r == size || nums[r]-nums[l] > 2*k {
				break
			}
			res = append(res, (nums[l]+nums[r])/2)
			used[l] = true
			used[r] = true
		}
		if len(res) == size/2 {
			return res
		}
	}
	return []int{}
}

func recoverArray1(nums []int) []int {
	sort.Ints(nums)
	size := len(nums)
	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[0]
		if d == 0 || d%2 == 1 {
			continue
		}
		var res []int
		res = append(res, (nums[0]+nums[i])/2)
		used := make([]bool, len(nums))
		used[i] = true
		used[0] = true
		k := d / 2
		for l, r := 1, i+1; r < len(nums); {
			if l < size && used[l] {
				l += 1
				continue
			}
			if r < size && (nums[r]-nums[l] < 2*k || used[r]) {
				r += 1
				continue
			}
			if r == size || nums[r]-nums[l] > 2*k {
				break
			}
			res = append(res, (nums[l]+nums[r])/2)
			used[l] = true
			used[r] = true
			l, r = l+1, r+1
		}
		if len(res) == size/2 {
			return res
		}
	}
	return []int{}
}

func main() {
	nums := []int{11, 6, 3, 4, 8, 7, 8, 7, 9, 8, 9, 10, 10, 2, 1, 9}
	fmt.Println(recoverArray1(nums))
}
