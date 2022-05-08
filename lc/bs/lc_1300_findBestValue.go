package main

import (
	"fmt"
	"sort"
)

/*
给你一个整数数组arr 和一个目标值target ，请你返回一个整数value，使得将数组中所有大于value 的值变成value 后，
数组的和最接近 target（最接近表示两者之差的绝对值最小）。
如果有多种使得和最接近target的方案，请你返回这些整数中的最小值。
请注意，答案不一定是arr 中的数字。

1、通过给出一个value，使得数组中所有>value的元素均变成value
2、返回一个value值，使得数组的元素和最接近target

1、数组排序
2、遍历数组，计算 tmp=pre + (n-i)*arr[i]，pre表示前i-1个数的和，上面式子表示将arr[i]替换为value
若value=arr[i]时 tmp<target，value=arr[i+1]时，tmp>=target，所以value一定在arr[i]与arr[i+1]之间
3、所以当遍历到arr[i]，且tmp>=target时，通过target和pre反推出一个可选value，计算value和value+1分别作为结果时，tmp和target的差值，
谁最接近，就返回对应的value
*/
func findBestValue(arr []int, target int) int {
	n, pre := len(arr), 0
	sort.Ints(arr)
	for i := 0; i < n; i++ {
		tmp := pre + (n-i)*arr[i]
		if tmp >= target {
			value := (target - pre) / (n - i)
			sumLow := pre + value*(n-i)
			sumHigh := pre + (value+1)*(n-i)
			if target-sumLow <= sumHigh-target {
				return value
			} else {
				return value + 1
			}
		}
		pre += arr[i]
	}
	return arr[len(arr)-1]
}

func main() {
	arr := []int{2, 3, 5}
	target := 10
	fmt.Println(findBestValue(arr, target))
}
