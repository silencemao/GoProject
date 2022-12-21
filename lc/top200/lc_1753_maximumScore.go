package main

import (
	"fmt"
	"sort"
)

/*

移除石子的最大得分
你正在玩一个单人游戏，面前放置着大小分别为 a、b、c的 三堆 石子。

每回合你都要从两个 不同的非空堆 中取出一颗石子，并在得分上加 1 分。当存在 两个或更多 的空堆时，游戏停止。

给你三个整数 a 、b 和 c ，返回可以得到的 最大分数 。
*/
func maximumScore(a int, b int, c int) int {
	tmp := []int{a, b, c}
	sort.Ints(tmp)
	if tmp[0]+tmp[1] <= tmp[2] {
		return tmp[0] + tmp[1]
	}
	return (tmp[0] + tmp[1] + tmp[2]) / 2
}

func maximumScore1(a int, b int, c int) int {
	res := 0
	tmp := []int{a, b, c}
	sort.Ints(tmp)
	for tmp[1] > 0 {
		res += 1
		tmp[2] -= 1
		tmp[1] -= 1
		sort.Ints(tmp)
	}
	return res
}

func main() {
	a, b, c := 4, 4, 6
	fmt.Println(maximumScore(a, b, c))
	fmt.Println(maximumScore1(a, b, c))
}
