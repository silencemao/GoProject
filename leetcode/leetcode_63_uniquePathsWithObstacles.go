package main

import "fmt"

/*
同62题，只是在矩阵中加入了障碍物，矩阵中的值为0表示可以通过，值为1表示为障碍物，不能通过。
对第一行和第一列单独处理时，要考虑到当前位置是否可以通过取决于前一个位置是否可以通过 即dp[i][0] = dp[i-1][0]
对于其它行，同第62题，只是要加一个判断当前位置是否为1
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0]==1 {
		return 0
	}

	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0]==0 {
			dp[i][0] = dp[i - 1][0]
		}
	}

	for j := 1; j < n; j++ {
		if obstacleGrid[0][j]==0 {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j]==1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]

}

func main() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
