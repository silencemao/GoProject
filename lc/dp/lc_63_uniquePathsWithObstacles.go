package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	if obstacleGrid[0][0] == 1 {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 { // 当前位置不是障碍物
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 { // 当前位置不是障碍物
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 { // 障碍物 当前位置无法到达
				dp[i][j] = 0
			} else {
				if dp[i-1][j] == 0 && dp[i][j-1] == 0 { // dp[i][j]==0表示当前位置是障碍物，无法到达
					dp[i][j] = 0
				} else if dp[i-1][j] == 0 {
					dp[i][j] = dp[i][j-1]
				} else if dp[i][j-1] == 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	if obstacleGrid[0][0] == 1 {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				if obstacleGrid[i-1][j] == 1 && obstacleGrid[i][j-1] == 1 {
					dp[i][j] = 0
				} else if obstacleGrid[i-1][j] == 1 || obstacleGrid[i][j-1] == 1 {
					dp[i][j] = util.MaxInt(dp[i-1][j], dp[i][j-1])
				} else {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
	fmt.Println(uniquePathsWithObstacles1(obstacleGrid))
}
