package main

import "fmt"

/*
给定两个数组 gas cost
gas表示在i位置可以加的油量，cost表示从i到i+1消耗的油量
判断是否可以从某个位置出发，返回到该位置，若可以返回 出发的位置 若不可以 返回-1

1、首先找到合适的起点(gas[i]>cost[i]) curOil = gas[i]
2、从当前位置开始向后走，满足当前油量 足够到达下一个位置 curOil > cost[i]
3、到达下一个位置之后加油 curOil += gas[i] - cost[i-1] , 此处注意 从起点出发时不需要减去消耗的油
4、继续2

*/
func canCompleteCircuit(gas []int, cost []int) int {
	res := -1
	var start, inds []int

	for i := 0; i < len(gas); i++ {
		if gas[i] >= cost[i] {
			start = append(start, i)
		}
		inds = append(inds, i)
	}

	for _, s := range start {
		tmp := make([]int, len(gas))
		copy(tmp, append(inds[s:], inds[:s]...))
		tmp = append(tmp, s) //回到起点
		curOil := 0
		for i := 0; i < len(tmp); i++ {
			if i == 0 { // 从起点出发，加油
				curOil += gas[tmp[i]]
			} else { //加油 减去从上一个地方到当前位置耗费的油
				curOil += gas[tmp[i]] - cost[tmp[i-1]]
			}
			if curOil < cost[tmp[i]] { // 当前的油是否够到下一个地方
				curOil = -1
				break
			}
		}
		if curOil >= 0 {
			return s
		}
	}
	return res
}

// 与上一步的做法一致，只是取消了ind的拷贝
func canCompleteCircuit1(gas []int, cost []int) int {
	res := -1
	for s := 0; s < len(gas); s++ {
		if gas[s] < cost[s] {
			continue
		}
		curOil := 0
		curOil += gas[s]
		for i := s + 1; i < len(gas); i++ {
			curOil += gas[i] - cost[i-1]
			if curOil < cost[i] {
				curOil = -1
				break
			}
		}
		if curOil < 0 {
			continue
		}

		for i := 0; i <= s; i++ {
			if i == 0 {
				curOil += gas[0] - cost[len(cost)-1]
			} else {
				curOil += gas[i] - cost[i-1]
			}
			if curOil < cost[i] {
				curOil = -1
				break
			}
		}
		if curOil >= 0 {
			return s
		}
	}
	return res
}

// 与上一步的做法一致，索引idx = (idx+1)%len(gas)
func canCompleteCircuit2(gas []int, cost []int) int {
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

// https://leetcode-cn.com/problems/gas-station/solution/gong-shui-san-xie-noxiang-xin-ke-xue-xi-zsgqp/
func canCompleteCircuit3(gas []int, cost []int) int {
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

func main() {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
	fmt.Println(canCompleteCircuit1(gas, cost))
	fmt.Println(canCompleteCircuit2(gas, cost))
	fmt.Println(canCompleteCircuit3(gas, cost))
}
