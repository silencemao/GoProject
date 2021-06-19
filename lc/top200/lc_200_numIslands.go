package main

import "fmt"

/*
岛屿的数量，给定一个二维数组，数组有0 1组成，1表示岛屿，0表示水。
岛屿与其上下左右的岛屿是联通的（表示一个岛屿），每个岛屿的四周都是被水包围着
返回岛屿的个数

		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
以上面的矩阵为例子，左上角的4个1构成一个岛屿，中间的一个1构成一个岛屿 右下角的两个1构成一个岛屿

我们可以逐个遍历，碰到1就计数加一，但是一个岛屿可能是多个1构成的，若是直接加一会重复计数，
因此我们需要在碰到一个岛屿的时候，向它的四周扩散，将与其联通的岛屿找出来，并标记，这样下次遍历到就不会重复计数了

1、找到一个岛屿 （grid[i][j]=1）
2、四周扩散 （上下左右）
3、标记（1->2 下次碰到就不重复计数了）
*/
func isIsland(grid [][]byte, i, j int) {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
		if grid[i][j] == '1' {
			grid[i][j] = '2'
			isIsland(grid, i-1, j)
			isIsland(grid, i+1, j)
			isIsland(grid, i, j-1)
			isIsland(grid, i, j+1)
		}
	}
}

func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '1' {
				continue
			}
			res += 1
			isIsland(grid, i, j)
		}
	}
	return res
}
func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
