package main

import "fmt"

/*
结合第207题，返回学习课程的顺序
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	res := []int{}

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

	for len(queue) > 0 { // 此处用队列是精髓
		first := queue[0]
		queue = queue[1:]
		res = append(res, first)

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
			return []int{}
		}
	}
	return res
}
func main() {
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	numCourses := 4

	fmt.Println(findOrder(numCourses, prerequisites))
}
