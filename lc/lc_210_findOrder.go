package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 1 {
		return []int{0}
	}

	graph := make(map[int][]int, len(prerequisites))
	in := make([]int, numCourses)
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
		in[pre[0]]++
	}

	var queue []int
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}
	var res []int
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		res = append(res, front)
		for _, next := range graph[front] {
			in[next]--
			if in[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if len(res) == numCourses {
		return res
	} else {
		return []int{}
	}
}

func findOrder1(numCourses int, prerequisites [][]int) []int {
	if numCourses == 1 {
		return []int{0}
	}

	graph := make(map[int][]int, len(prerequisites))
	for _, pre := range prerequisites {
		graph[pre[0]] = append(graph[pre[0]], pre[1])
	}
	visited := make([]int, numCourses)
	var res []int
	for i := 0; i < numCourses; i++ {
		if !dfsFindOrder(graph, &visited, i, &res) {
			return []int{}
		}
	}
	return res
}

func dfsFindOrder(graph map[int][]int, visited *[]int, i int, res *[]int) bool {
	if (*visited)[i] == 0 {
		(*visited)[i] = 1
		for k := 0; k < len(graph[i]); k++ {
			if !dfsFindOrder(graph, visited, graph[i][k], res) {
				return false
			}
		}
		(*visited)[i] = 2
	} else if (*visited)[i] == 1 {
		return false
	} else {
		return true
	}
	*res = append(*res, i)
	return true
}

func main() {
	numCourse := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Println(findOrder(numCourse, prerequisites))
	fmt.Println(findOrder1(numCourse, prerequisites))
}
