package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
给你一个二进制字符串数组 strs 和两个整数 m 和 n，strs中的字符串由0和1组成
请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。
如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
找出最大子集，子集中的元素0最多由m个，1最多有n个。

此题可以转换为背包问题来解决，m，n就是背包的容量(之前背包容量是一个数字，现在是两个数字了)，strs就是填充到背包的物品，需要对物品进行拆解一下，
物品拆成x个0，y个1，填充到背包中。

1、首先我们定义一个二维数组，dp[m+1][n+1] dp[0][0]=0 表示组成0个0，0个1需要0个字符串
2、因为每个字符串中的0和1的个数经常要用到，为了方便，所以可以先统计出每个字符串中0和1 的个数，存到一个map中
3、开始往背包中填充物品，
遍历字符串，拿到字符串中0和1的个数
从大到小遍历m，n，若此时的m，n大于等于字符串中0、1的个数，可以考虑放入，比较dp[i-cnt0][cnt1]+1 dp[i][j]的大小

*/
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	strCnt := make(map[string][]int, len(strs))
	for _, str := range strs {
		strCnt[str] = make([]int, 2)
		for _, b := range str {
			strCnt[str][b-'0']++
		}
	}

	for _, str := range strs {
		cnt := strCnt[str]
		for i := m; i >= 0 && i >= cnt[0]; i-- {
			for j := n; j >= 0 && j >= cnt[1]; j-- {
				dp[i][j] = util.MaxInt(dp[i-cnt[0]][j-cnt[1]]+1, dp[i][j])
			}
		}
	}

	return dp[m][n]
}
func main() {
	strs := []string{"10", "1", "0"}
	m, n := 1, 1
	fmt.Println(findMaxForm(strs, m, n))
}
