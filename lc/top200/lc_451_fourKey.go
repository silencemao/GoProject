package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给定一个键盘，只有4个键，A， ctrl-C，ctrl-A，ctrl-V
一共有N次按键机会，最终屏幕上会有多少个A
*/
func help451(N int, aNum int, copy int) int {
	if N <= 0 {
		return aNum
	}
	return util.MaxInt(util.MaxInt(help451(N-2, aNum, aNum), help451(N-1, aNum+1, copy)), help451(N-1, aNum+copy, copy))
}
func maxA(N int) int {

	return help451(N, 0, 0)
}

func maxA1(N int) int {
	dp := make([]int, N+1)
	for i := 1; i <= N; i++ {
		dp[i] = dp[i-1] + 1 // 直接an一个A
		for j := 2; j < i; j++ {
			// crtl-A, crtl-C dp[j-2]处的A，复制i-j次，则屏幕上共有dp[j-2]*(i-j+1)个A
			dp[i] = util.MaxInt(dp[i], dp[j-2]*(i-j+1))
		}
	}
	return dp[N]
}

func main() {
	N := 7
	fmt.Println(maxA(N))
	fmt.Println(maxA1(N))
}
