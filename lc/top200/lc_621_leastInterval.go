package main

import "fmt"

/*
任务调度器

给你一个用字符数组tasks 表示的 CPU 需要执行的任务列表。其中每个字母表示一种不同种类的任务。
任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。在任何一个单位时间，CPU 可以完成一个任务，或者处于待命状态。
然而，两个 相同种类 的任务之间必须有长度为整数 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
你需要计算完成所有任务所需要的 最短时间 。

tasks中每个字母表示一种类型的任务，每个任务的执行花费1个单位时间，相同任务的执行中间必须有n个时间单位的间隔。
完成tasks中的任务至少要有多少个时间单位

输入：tasks = ["A","A","A","B","B","B"], n = 2    输出：8
解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B
     在本示例中，两个相同类型任务之间必须间隔长度为 n = 2 的冷却时间，而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。

1、找出同类型任务数量最多的任务数 max
2、找出任务数量为max的任务种类数 maxCnt
3、计算总时间 (max-1)*(n+1) + maxCnt
max-1表示 任务数量，对于出现次数最多的任务种类A，两个任务A之间需要间隔n个单位，所以执行完一个任务A之后，需要n个单位间隔才能执行下一个任务A，
即两个任务A之间的时间间隔是n+1
*/
func leastInterval(tasks []byte, n int) int {
	res := len(tasks)

	set := make([]int, 26)
	for _, b := range tasks {
		set[b-'A'] += 1
	}
	maxCnt, max := 0, 0
	for _, num := range set {
		if num > max {
			max = num
			maxCnt = 1
		} else if num == max {
			maxCnt += 1
		}
	}
	if res < (max-1)*(n+1)+maxCnt {
		res = (max-1)*(n+1) + maxCnt
	}

	return res
}

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	fmt.Println(leastInterval(tasks, n))
}
