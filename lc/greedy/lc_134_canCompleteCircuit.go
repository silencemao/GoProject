package main

import "fmt"

/*
在一条环路上有N个加油站，其中第i个加油站有汽油gas[i]升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

说明:

如果题目有解，该答案即为唯一答案。
输入数组均为非空数组，且长度相同。
输入数组中的元素均为非负数。

*/
// 超时
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	for i := 0; i < n; i++ {
		if gas[i] >= cost[i] {

			cur := gas[i]
			pre, next := i, (i+1)%n
			for i != next {
				cur = cur - cost[pre] + gas[next]
				if cur < cost[next] {
					break
				}
				pre, next = next, (next+1)%n
			}
			if next == i {
				return next
			}
		}
	}
	return -1
}

// 超时
func canCompleteCircuit2(gas []int, cost []int) int {
	n := len(gas)

	for i := 0; i < n; i++ {
		if gas[i] >= cost[i] {
			cur := gas[i] - cost[i]
			next := (i + 1) % n
			for i != next {
				cur += gas[next] - cost[next]
				if cur < 0 {
					break
				}
				next = (next + 1) % n
			}
			if next == i {
				return next
			}
		}
	}
	return -1
}

// 超时
func canCompleteCircuit3(gas []int, cost []int) int {
	res := -1
	n := len(gas)
	for s := 0; s < len(gas); s++ {
		if gas[s] < cost[s] {
			continue
		}
		curOil := gas[s] - cost[s]
		idx := (s + 1) % n
		for idx != s {
			curOil += gas[idx] - cost[idx]
			if curOil < 0 {
				break
			}
			idx = (idx + 1) % n
		}
		if idx == s {
			return s
		}
	}
	return res
}

// 通过
func canCompleteCircuit4(gas []int, cost []int) int {
	res := -1
	n := len(gas)
	for s := 0; s < n; {
		if gas[s] < cost[s] {
			s++
			continue
		}
		curOil := gas[s] - cost[s]
		idx := s + 1 // idx 一直累加，否则idx会一直 < n 陷入死循环
		for idx%n != s {
			curOil += gas[idx%n] - cost[idx%n]
			if curOil < 0 {
				break
			}
			idx = idx + 1
		}
		if idx%n == s {
			return s
		} else { // 从s出发 到了idx 油量为负，直接从idx作为新起点
			s = idx
		}
	}
	return res
}

// 通过
func canCompleteCircuit5(gas []int, cost []int) int {
	curSum, totalSum := 0, 0
	start := 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1 // [j, i]之间剩余油量<0,这之间的数都不能作为起点了 很关键
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}

func main() {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	fmt.Println(canCompleteCircuit2(gas, cost))
}
