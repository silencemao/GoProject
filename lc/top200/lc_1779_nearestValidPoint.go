package main

import "GoProject/leetcode/util"

/*
找到最近的有相同 X 或 Y 坐标的点
*/
func nearestValidPoint(x int, y int, points [][]int) int {
	ind, dis := -1, 1<<31-1
	for i, p := range points {
		if p[0] == x || p[1] == y {
			tmpDis := disManhadun([]int{x, y}, p)
			if tmpDis < dis {
				dis = tmpDis
				ind = i
			}
		}
	}
	return ind
}

func disManhadun(p1, p2 []int) int {
	return util.Abs(p1[0]-p2[0]) + util.Abs(p1[1]-p2[1])
}

func main() {

}
