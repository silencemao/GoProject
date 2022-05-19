package main

import (
	"fmt"
	"sort"
)

/*
在社交媒体网站上有 n 个用户。给你一个整数数组 ages ，其中 ages[i] 是第 i 个用户的年龄。
如果下述任意一个条件为真，那么用户 x 将不会向用户 y（x != y）发送好友请求：

ages[y] <= 0.5 * ages[x] + 7
ages[y] > ages[x]
ages[y] > 100 && ages[x] < 100
*/
func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	res, n := 0, len(ages)
	j, k := 0, 1 //给y发消息的最小x索引，给y发消息的最大x索引
	for i := 0; i < n; i++ {
		y := ages[i]
		for j < i && !help825(ages[j], y) {
			j += 1
		}
		if k < i {
			k = i
		}
		for k < n && help825(ages[k], y) {
			k += 1
		}
		if k > j {
			res += k - j - 1 // 不包括i和k
		}
	}
	return res
}

func help825(x, y int) bool {
	if (y-7)*2 <= x || y > x || (y > 100 && x < 100) {
		return false
	}
	return true
}

func main() {
	ages := []int{118, 14, 7, 63, 103}
	fmt.Println(numFriendRequests(ages))
}
