package main

import "fmt"

/*
我们称一个分割整数数组的方案是 好的，当它满足：

数组被分成三个 非空连续子数组，从左至右分别命名为left，mid，right。
left中元素和小于等于mid中元素和，mid中元素和小于等于right中元素和。
给你一个 非负 整数数组nums，请你返回好的 分割 nums方案数目。由于答案可能会很大，请你将结果对 109+ 7取余后返回。
1、前缀和
2、先确定第一刀的切分位置，然后二分法确定第二刀切分的左右边界 （二分法，左边界 第一个>=target1，最后一个<=target1）
*/
func waysToSplit(nums []int) int {
	n := len(nums)
	sums := make([]int, n+1)
	res, mod := int64(0), 1e9+7

	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	for i := 1; i < n; i++ { // left 和 mid两部分的分界点

		if sums[i]*3 > sums[n] {
			continue
		}

		l, r := i+1, n-1 // mid和right两部分分界点的最右边界
		for l <= r {
			mid := l + (r-l)>>1
			if sums[n]-sums[mid] >= sums[mid]-sums[i] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		ll, rr := i+1, n-1
		for ll <= rr { // mid和right两部分分界点的最左边界
			mid := ll + (rr-ll)>>1
			if sums[mid]-sums[i] < sums[i] {
				ll = mid + 1
			} else {
				rr = mid - 1
			}
		}
		res += int64(r - ll + 1)
	}

	return int(res % int64(mod))
}

/*
target
l:找第一个>=sums[i]*2的位置
r:找最后一个 满足 sums[n]-sums[mid] >= sums[mid]-sums[i]  ==> (sums[n]+sums[i]) / 2 >= sums[mid]
即找到最后一个<=(sums[n]+sums[i]) / 2
*/
func waysToSplit1(nums []int) int {
	n := len(nums)
	sums := make([]int, n+1)
	res, mod := int64(0), 1e9+7

	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	for i := 1; i < n; i++ { // left 和 mid两部分的分界点

		if sums[i]*3 > sums[n] {
			continue
		}

		l, r, target := i+1, n, (sums[n]+sums[i])/2 // 找最后一个小于等于target的索引
		for l < r {                                 //mid和right两部分分界点的最右边界
			mid := l + (r-l)>>1
			if sums[mid] <= target {
				if mid+1 == len(sums)-1 || sums[mid+1] > target {
					r = mid
					break
				}
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		ll, rr, target := i+1, n, 2*sums[i] // 找第一个大于等于target的索引
		for ll < rr {                       // mid和right两部分分界点的最左边界
			mid := ll + (rr-ll)>>1
			if sums[mid] >= target {
				if mid-1 == i+1 || sums[mid-1] < target {
					ll = mid
					break
				}
				rr = mid - 1
			} else {
				ll = mid + 1
			}
		}
		if r >= ll {
			res += int64(r - ll + 1)
		}
	}

	return int(res % int64(mod))
}

func main() {
	nums := []int{1, 2, 2, 2, 5, 0}
	fmt.Println(waysToSplit(nums))
	fmt.Println(waysToSplit1(nums))
}
