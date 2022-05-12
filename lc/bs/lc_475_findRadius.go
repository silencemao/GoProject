package main

import (
	"GoProject/leetcode/util"
	"fmt"
	"sort"
)

/*
冬季已经来临。你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
在加热器的加热半径范围内的每个房屋都可以获得供暖。
现在，给出位于一条水平线上的房屋houses 和供暖器heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。
说明：所有供暖器都遵循你的半径标准，加热的半径也一样。
*/
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	res := 0
	for i, j := 0, 0; i < len(houses); i++ {
		cur := util.Abs(houses[i] - heaters[j])
		for j < len(heaters) && heaters[j] < houses[i] {
			cur = houses[i] - heaters[j]
			j++
		}
		if j < len(heaters) {
			if v := heaters[j] - houses[i]; v < cur {
				cur = v
			}
		}
		if cur > res {
			res = cur
		}
		if j > 0 {
			j -= 1
		}
	}
	return res
}

func main() {
	houses := []int{1, 2, 3, 4}
	heaters := []int{1, 4}
	fmt.Println(findRadius(houses, heaters))

}
