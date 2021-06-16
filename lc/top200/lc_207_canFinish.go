package main

import "fmt"

/*
你这个学期必须选修 numCourses 门课程，记为0到numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组prerequisites 给出，其中prerequisites[i] = [ai, bi] ，
表示如果要学习课程ai 则 必须 先学习课程 bi 。

例如，先修课程对[0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
https://leetcode-cn.com/problems/course-schedule

*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	courses := make([]int, numCourses)
	relation := make(map[int][]int, len(prerequisites))

	for _, p := range prerequisites {
		courses[p[0]] += 1                            // 每门课需要依赖几门课
		relation[p[1]] = append(relation[p[1]], p[0]) // 学完这门可以学哪些
	}

	queue := []int{}
	for i, v := range courses { // 可以直接学的课程
		if v == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		if nexts, ok := relation[first]; ok {
			for _, next := range nexts {
				courses[next] -= 1
				if courses[next] == 0 {
					queue = append(queue, next)
				}
			}
		}
	}

	for _, v := range courses {
		if v > 0 {
			return false
		}
	}

	return true
}
func main() {
	numCourses := 5
	prerequisites := [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}
	fmt.Println(canFinish(numCourses, prerequisites))
}
