package main

import (
	"fmt"
	"sort"
)

/*
https://leetcode.com/problems/two-city-scheduling/discuss/311050/Go-for-sort
*/
type byCost [][]int

func (c byCost) Len() int {
	return len(c)
}

func (c byCost) Swap(i int, j int) {
	c[i][0], c[j][0] = c[j][0], c[i][0]
	c[i][1], c[j][1] = c[j][1], c[i][1]
}

func (c byCost) Less(i int, j int) bool {
	return (c[i][0] - c[i][1]) < (c[j][0] - c[j][1])
}

/*
给定一个二维数组costs。
有A B两个城市，有N个人，每个人去A城市和去B城市的费用为costs[i][0], costs[i][1]
如何分配这N个人，N/2个人去A城市，N/2个人去B城市，并使得总费用最小

一半人去A城市，一半人去B城市，总费用最小
如何选择去A城市而不去B城市的这部分人，也就是说这一部分人去A城市的费用要比去B城市的费用少，所以可以按照(costs[i][0] - costs[i][1])来从小
到大排序，即前一半的人去A城市比去B城市的费用之和要小，即后一半的人去B城市的费用要比去A城市的费用之和要小，所以总费用是最小的。
*/
func twoCitySchedCost(cost [][]int) int {
	res := 0

	sort.Sort(byCost(cost))
	for i := 0; i < len(cost) / 2; i++ {
		res += cost[i][0]
		res += cost[len(cost) - i - 1][1]
	}

	return res
}

func main() {
	inputAry := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}
	fmt.Println(twoCitySchedCost(inputAry))
}
