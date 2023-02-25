package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
给出一部分员工的名字keyName，以及与该名字对应的时间，表示每个员工的打卡时间keyTime，若一个员工在60min内，打卡三次及以上，则表示违规，
找出所有违规的员工
*/
func alertName(keyName, keyTime []string) []string {
	var res []string
	set := make(map[string][]int)
	for i, t := range keyTime {
		name := keyName[i]
		h, _ := strconv.Atoi(t[:2])
		m, _ := strconv.Atoi(t[3:])
		set[name] = append(set[name], h*60+m)
	}
	for name, times := range set {
		sort.Ints(times)
		fmt.Println(name, times)
		for i, t := range times[2:] {
			if t-times[i] <= 60 {
				res = append(res, name)
				break
			}
		}
	}
	sort.Strings(res)
	return res
}

func main() {
	keyName := []string{"a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "c", "c", "c", "c"}
	ketTime := []string{"01:35", "08:43", "20:49", "00:01", "17:44", "02:50", "18:48", "22:27", "14:12", "18:00", "12:38", "20:40", "03:59", "22:24"}
	fmt.Println(alertName(keyName, ketTime))
}
