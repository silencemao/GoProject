package main

import (
	"fmt"
	"GoProject/leetcode/util"
)

/*
给定两个字符串 word1 word2，可以对word1中每一个字母进行替换、增加、删除操作，使word1变成word2需要的最少步数
采用动态规划的思路，生成一个m+1 * n+1大小的二维数组，第0行，都0列表示将空字符串""转换为某一个个字串需要的步数
当word1[i-1]==word2[j-1]时，直接将res[i-1][j-1]处的值覆盖到当前位置处res[i][j]处，表示当前位置不需要做任何改变，
在上一步做完该表就能满足word1的子串与word2的子串相等
若word1[i-1]!=word2[j-1]，说吗当前位置需要作出调整，只需要在当前位置的左上位置，左边位置、上方位置处三者最小的值上加一即可，
即在最小变动的基础上再变一次就能满足
*/

func minDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	res := make([][]int, m+1)
	for i := range res {
		res[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		res[i][0] = i
	}

	for j := 1; j < n+1; j++ {
		res[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				res[i][j] = res[i-1][j-1]
			} else {
				res[i][j] = util.MinInt(res[i-1][j-1], util.MinInt(res[i][j-1], res[i-1][j])) + 1
			}
		}
	}
	return res[m][n]
}

/*
*/
func minDistance1(word1, word2 string) int {
	m, n := len(word1), len(word2)
	res := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		res[i] = i
	}
	for i := 1; i < m+1; i++ {
		pre := res[0]           // res[i-1][j-1]
		res[0] = i
		for j := 1; j < n+1; j++ {
			temp := res[j]      // 此处res[i-1][j]
			if word1[i-1] == word2[j-1] {
				res[j] = pre
			} else {
				// 此处res[j-1]相当于 res[i-1][j-1] res[j]相当于res[i-1][j]
				res[j] = util.MinInt(pre, util.MinInt(res[j], res[j-1])) + 1
			}
			pre = temp          // 更新pre
		}
	}
	return res[n]
}

func main() {
	word1, word2 := "intention", "execution"
	fmt.Println(minDistance(word1, word2))
	fmt.Println(minDistance1(word1, word2))
}
