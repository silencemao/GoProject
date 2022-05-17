package main

import "fmt"

/*
符合下列属性的数组 arr 称为 山脉数组 ：
arr.length >= 3
存在 i（0 < i< arr.length - 1）使得：
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]
给你由整数组成的山脉数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i 。

*/
func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r + 1) >> 1 // 防止mid为0
		if arr[mid-1] < arr[mid] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func peakIndexInMountainArray1(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) >> 1 // mid<r
		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
func main() {
	arr := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	fmt.Println(peakIndexInMountainArray(arr))
	fmt.Println(peakIndexInMountainArray1(arr))
}
