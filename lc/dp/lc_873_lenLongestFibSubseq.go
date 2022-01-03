package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
最长的斐波那契子序列的长度
[1,2,3,4,5,6,7,8]

1、从索引为2开始遍历数组
2、如何判断arr[i]可以和前面的数组构成斐波那契数列
   arr[i]和arr[j] (j=0...i-1)做差得到v，若v存在与arr中，表示v,arr[j],arr[i]可以构成斐波那契子序列
   长度为以j结尾时，斐波那契子序列的长度+1
3、如何存储斐波那契子序列的长度，采用二维数组 dp[j][i]最后两项是arr[j], arr[i]构成的斐波那契子序列的长度，如第二步所示
4、下面判断k<j有两层含义，第一是表示倒数第二项为arr[j]，第二是若j<k时，dp[j][i]已经计算过了(有可能已经+1)，再次计算会重复计算并且出错
*/
func lenLongestFibSubseq(arr []int) int {
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}
	set := make(map[int]int, 0)
	for i, v := range arr {
		set[v] = i
	}

	res := 0
	for i := 2; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			v := arr[i] - arr[j]
			if k, ok := set[v]; ok { // arr中存在差值
				if k < j {
					dp[j][i] = util.MaxInt(dp[j][i], dp[k][j]+1)
					res = util.MaxInt(res, dp[j][i]+2)
				}
			}
		}
	}
	return res
}
func main() {
	arr := []int{2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50}
	fmt.Println(lenLongestFibSubseq(arr))
}
