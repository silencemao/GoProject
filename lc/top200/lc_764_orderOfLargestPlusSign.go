package main

import "GoProject/leetcode/util"

/*
在一个 n x n 的矩阵grid中，除了在数组mines中给出的元素为0，其他每个元素都为1。mines[i] = [xi, yi]表示grid[xi][yi] == 0

返回 grid 中包含1的最大的 轴对齐 加号标志的阶数 。如果未找到加号标志，则返回 0 。

一个k阶由1组成的 “轴对称”加号标志 具有中心网格grid[r][c] == 1，以及4个从中心向上、向下、向左、向右延伸，长度为k-1，由1组成的臂。注意，只有加号标志的所有网格要求为 1 ，别的网格可能为 0 也可能为 1 。
*/
func orderOfLargestPlusSign(n int, mines [][]int) (ans int) {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := 0; j <= n; j++ {
			dp[i][j] = n
		}
	}
	bans := map[int]bool{}
	for _, m := range mines {
		bans[m[0]*n+m[1]] = true
	}
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if bans[i*n+j] {
				count = 0
			} else {
				count += 1
			}
			dp[i][j] = util.MinInt(dp[i][j], count)
		}
		count = 0
		for j := n - 1; j >= 0; j-- {
			if bans[i*n+j] {
				count = 0
			} else {
				count += 1
			}
			dp[i][j] = util.MinInt(dp[i][j], count)
		}
	}
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if bans[j*n+i] {
				count = 0
			} else {
				count += 1
			}
			dp[j][i] = util.MinInt(dp[j][i], count)
		}
		count = 0
		for j := n - 1; j >= 0; j-- {
			if bans[j*n+i] {
				count = 0
			} else {
				count += 1
			}
			dp[j][i] = util.MinInt(dp[j][i], count)
			ans = util.MaxInt(dp[j][i], ans)
		}
	}
	return ans
}

func main() {

}
