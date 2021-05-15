package main

import "fmt"

/*
给定 numCourse个课，但是每个课的执行顺序是由前提的[1, 0]表示，要想上1这门课，必须先上0这门课
然后根据给定的课程数，以及每个课执行的顺序，来确定是否能完成这些课

这个题的本意是考察在有向图中是否有环，若有环则表明不能完成这些课，若无环则表明可以完成

首先建立图，即每个课执行完可以执行哪些课，同时对后面要执行的出度加一，
然后将出度为0的课放入队列，出度为0，即执行这个课不需要先执行其他课
执行完出度为0的课，其后面的课的出度就可以减一，然后再将出度为0的课加入队列中
最后队列中所有的课的出度都为0， 表示可以完成，若还有出度为1的，则表示不能完成
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 1 {
		return true
	}
	graph := make(map[int][]int, len(prerequisites))
	in := make([]int, numCourses)
	var queue []int
	for _, a := range prerequisites {
		graph[a[1]] = append(graph[a[1]], a[0])
		in[a[0]]++
	}
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		for _, next := range graph[front] {
			in[next]--
			if in[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	for _, pre := range in {
		if pre != 0 {
			return false
		}
	}
	return true
}

func main() {
	numCourses := 2
	prerequisite := [][]int{{1, 0}}
	fmt.Println(canFinish(numCourses, prerequisite))
}
