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

1、找每个房间最近的暖气，分别是距离house[i]最近的前后各一个暖气，计算二者最近的距离d=min(d1,d2)
2、选择所有房间最近距离d的最大值，max(d)
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

func findRadius1(houses []int, heaters []int) int {
	res := -1
	for i := 0; i < len(houses); i++ {
		tmp := 1<<31 - 1
		for j := 0; j < len(heaters); j++ {
			tmp = util.MinInt(util.Abs(houses[i]-heaters[j]), tmp)
		}
		res = util.MaxInt(tmp, res)
	}

	return res
}

func findRadius2(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	res := -1
	h1, h2 := heaters[0], heaters[len(heaters)-1]
	for i, house := range houses {
		if house <= h1 {
			res = util.MaxInt(h1-house, res)
		} else if house >= h2 {
			res = util.MaxInt(house-h2, res)
		} else {
			pos := bs475(heaters, house)
			tmp := util.MinInt(heaters[pos]-house, house-heaters[pos-1])
			res = util.MaxInt(tmp, res)
		}
		fmt.Println(res, i)
	}
	return res
}

// 找第一个>=target的值
func bs475(heaters []int, target int) int {
	l, r := 0, len(heaters)-1
	for l <= r {
		mid := l + (r-l)>>1
		if heaters[mid] >= target {
			if mid == 0 || heaters[mid-1] < target {
				return mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	houses := []int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923}
	heaters := []int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612}
	fmt.Println(findRadius(houses, heaters))
	fmt.Println(findRadius1(houses, heaters))
	fmt.Println(findRadius2(houses, heaters))
}
