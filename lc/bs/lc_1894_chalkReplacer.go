package main

import "fmt"

/*
一个班级里有n个学生，编号为 0到 n - 1。每个学生会依次回答问题，编号为 0的学生先回答，然后是编号为 1的学生，
以此类推，直到编号为 n - 1的学生，然后老师会重复这个过程，重新从编号为 0的学生开始回答问题。

给你一个长度为 n且下标从 0开始的整数数组chalk和一个整数k。一开始粉笔盒里总共有k支粉笔。当编号为i的学生回答问题时，他会消耗 chalk[i]支粉笔。
如果剩余粉笔数量 严格小于chalk[i]，那么学生 i需要 补充粉笔。

请你返回需要 补充粉笔的学生 编号。
*/
// 前缀和+二分法
func chalkReplacer(chalk []int, k int) int {
	sums := make([]int, len(chalk)+1)
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + chalk[i-1]
	}
	k = k % sums[len(sums)-1]
	l, r := 1, len(sums)-1
	for l < r {
		mid := l + (r-l)>>1
		if sums[mid] > k {
			r = mid
		} else {
			l = mid + 1 // 当前mid明确不需要补充粉笔，进行下一个
		}
	}
	return r - 1
}

// 变异前缀和+遍历
func chalkReplacer1(chalk []int, k int) int {
	sums := 0
	for _, num := range chalk {
		sums += num
	}
	k = k % sums
	for i := 0; i < len(chalk); i++ {
		k -= chalk[i]
		if k < 0 {
			return i
		}
	}
	return -1
}

func main() {
	chalk := []int{3, 4, 1, 2}
	k := 25
	fmt.Println(chalkReplacer(chalk, k))
}
