package main

import "fmt"

/*
寻找图中是否存在最短路径
有一个具有 n 个顶点的 双向 图，其中每个顶点标记从 0 到 n - 1（包含 0 和 n - 1）。
图中的边用一个二维整数数组 edges 表示，其中 edges[i] = [ui, vi] 表示顶点 ui 和顶点 vi 之间的双向边。
每个顶点对由 最多一条 边连接，并且没有顶点存在与自身相连的边。
请你确定是否存在从顶点 source 开始，到顶点 destination 结束的 有效路径 。
给你数组 edges 和整数 n、source 和 destination，如果从 source 到 destination 存在 有效路径 ，则返回 true，否则返回 false 。

*/
func validPath(n int, edges [][]int, source int, destination int) bool {
	vis := make([]bool, n)
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == destination {
			return true
		}
		vis[i] = true
		for _, j := range graph[i] {
			if !vis[j] && dfs(j) {
				return true
			}
		}
		return false
	}
	return dfs(source)
}

func validPath1(n int, edges [][]int, source int, destination int) bool {
	vis := make([]bool, n)
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	var bfs func(i int)
	bfs = func(i int) {
		var queue []int
		queue = append(queue, i)

		for len(queue) > 0 {
			front := queue[0]
			queue = queue[1:]
			vis[front] = true
			if front == destination {
				break
			}
			for _, j := range graph[front] {
				if !vis[j] {
					queue = append(queue, j)
					vis[j] = true
				}
			}
		}
	}
	bfs(source)
	return vis[destination]
}

func validPath2(n int, edges [][]int, source int, destination int) bool {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		x := i
		for x != p[x] {
			x = p[x]
		}

		for p[i] != x {
			p[i], i = x, p[i]
		}
		return p[x]
	}
	for _, edge := range edges {
		p[find(edge[0])] = find(edge[1])
	}

	return find(source) == find(destination)
}

func main() {
	n, source, destination := 6, 0, 5
	edges := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	fmt.Println(validPath(n, edges, source, destination))
	fmt.Println(validPath1(n, edges, source, destination))
	fmt.Println(validPath2(n, edges, source, destination))
}
