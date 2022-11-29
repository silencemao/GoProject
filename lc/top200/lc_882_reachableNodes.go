package main

import (
	"GoProject/leetcode/util"
	"container/heap"
	"fmt"
	"math"
)

func reachableNodes(edges [][]int, maxMoves, n int) (ans int) {
	g := make([][]neighbor, n)
	for _, e := range edges {
		u, v, cnt := e[0], e[1], e[2]
		g[u] = append(g[u], neighbor{v, cnt + 1})
		g[v] = append(g[v], neighbor{u, cnt + 1}) // 建图
	}

	dist := dijkstra(g, 0) // 从 0 出发的最短路

	for _, d := range dist {
		if d <= maxMoves { // 这个点可以在 maxMoves 步内到达
			ans++
		}
	}
	for _, e := range edges {
		u, v, cnt := e[0], e[1], e[2]
		a := util.MaxInt(maxMoves-dist[u], 0)
		b := util.MaxInt(maxMoves-dist[v], 0)
		ans += util.MinInt(a+b, cnt) // 这条边上可以到达的节点数
	}
	return
}

// 以下为 Dijkstra 算法模板
type neighbor struct{ to, weight int }

func dijkstra(g [][]neighbor, start int) []int {
	dist := make([]int, len(g))
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0
	h := hp8{{start, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair882)
		x := p.x
		if dist[x] < p.dist {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			if d := dist[x] + e.weight; d < dist[y] {
				dist[y] = d
				heap.Push(&h, pair882{y, d})
			}
		}
	}
	return dist
}

type pair882 struct{ x, dist int }
type hp8 []pair882

func (h hp8) Len() int            { return len(h) }
func (h hp8) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h hp8) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp8) Push(v interface{}) { *h = append(*h, v.(pair882)) }

//func (h *hp8) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp8) Pop() (v interface{}) { a := *h; *h, v = a[1:], a[0]; return }

func main() {
	edges := [][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}}
	maxMoves := 6
	n := 3
	fmt.Println(reachableNodes(edges, maxMoves, n))
}
