package main

import "fmt"

/*
给定一个二维数组，数组从左向右逐渐递增，从上到下逐渐递增。找出数组中第k小的数字
*/
func kthSmallest(nums [][]int, k int) int {
	n := len(nums)
	l, h := nums[0][0], nums[n-1][n-1]
	for l < h {
		mid := (l + h) / 2
		cnt := getCount(nums, mid)
		if cnt < k { // 说明目标值在mid右侧，不包括mid
			l = mid + 1
		} else { // cnt>=k 说明目标值在mid左侧，包括mid
			h = mid
		}
	}
	return h
}

func getCount(nums [][]int, mid int) int {
	cnt := 0
	n := len(nums)
	for j := 0; j < n; j++ {
		for i := n - 1; i >= 0; i-- {
			if nums[i][j] <= mid {
				cnt += i + 1
				break
			}
		}
	}
	return cnt
}

/*
给定二维数组，第k大的数字如何计算
将第k大转换为
第n*n-k+1小的数字
*/

func main() {
	nums := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	k := 7
	fmt.Println(kthSmallest(nums, k))
}

//l=1 h=15 mid=8 cnt=2 l=9
//9 15 mid=12 cnt=6 l=10
//10 15 mid 12 cnt=6 l=13
//13 15 mid 14 cnt=8 l=13
//13 13 mid 13 cnt=8
/*
一开始不断得通过left和right不断的夹逼，一定会使得left或right有一个先等于target的。如果left就是target(第k小的数)，此时righ肯定大于等于target，
以后不断取中值center=(left+right也)/2，既可以满足比center小的数有k个，也可以不断的缩小right的值，最终使得left=right=target，
这时候的center=target，从而找到第k小的数。（例如：1 2 4 7 11 23，target=4，就可以体现上面的过程）
*/
