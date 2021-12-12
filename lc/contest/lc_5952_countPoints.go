package main

import "fmt"

/*
环扣杆
https://leetcode-cn.com/problems/rings-and-rods/
有10个杆 3中颜色 分别为RGB
rings = "B0B6G0R6R0R6G9" 字符串长度为2n，表示n个环在杆上的分布，以B0为例，表示第0个杆上有一个颜色为B的环，G0表示第0个杆上有一个颜色为G的环

统计每个杆上的环，返回有多少个杆 是包含RGB三种颜色的环的，
*/
func countPoints(rings string) int {
	trans := make(map[byte]int)
	trans['R'] = 0
	trans['G'] = 1
	trans['B'] = 2
	set := make(map[byte][3]int, 0)
	rs := []byte(rings)
	for i := 1; i < len(rs); i += 2 {
		c := trans[rs[i-1]]
		n := rs[i] - '0'
		if ns, ok := set[n]; ok {
			ns[c] += 1
			set[n] = ns
		} else {
			ns[c] = 1
			set[n] = ns
		}
	}

	res := 0
	for _, vs := range set {
		isOk := true
		for _, v := range vs {
			if v <= 0 {
				isOk = false
				break
			}
		}
		if isOk {
			res += 1
		}
	}
	return res
}

func countPoints1(rings string) int {
	res := 0
	status := make([]int, 10) // 共计10个杆
	for i := 1; i < len(rings); i += 2 {
		c := rings[i-1]
		ind := rings[i] - '0'
		if c == 'R' {
			status[ind] |= 1
		} else if c == 'G' {
			status[ind] |= 2
		} else if c == 'B' {
			status[ind] |= 4
		}
	}
	for i := range status {
		if status[i] == 7 {
			res += 1
		}
	}
	return res
}

func main() {
	rings := "B0B6G0R6R0R6G9"
	fmt.Println(countPoints(rings))
	fmt.Println(countPoints1(rings))
}
