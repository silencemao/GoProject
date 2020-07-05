package main

import "fmt"

// https://github.com/muzixing/graph_algorithm/blob/master/dijkstra.py
//https://silencemao.github.io/2020/07/05/dijkstra%E7%AE%97%E6%B3%95/#more

const(
	MaxDis  int = 1<<7-1
)

type Dijkstra struct {
	tPints  []string
	tTwoPointDis map[string]int
}

func (d *Dijkstra) Init(tPoints []string, tDis [][]int) {
	if len(tPoints) != len(tDis) {
		panic("点数与矩阵的大小不一致")
	}
	d.tTwoPointDis = make(map[string]int, 0)
	for i := 0; i < len(tPoints); i++ {
		for j := 0; j < len(tPoints); j++ {
			key := tPoints[i] + "_" + tPoints[j]
			d.tTwoPointDis[key] = tDis[i][j]
		}
	}
	d.tPints = tPoints
}

func (d *Dijkstra) dijkstra() {
	tPoints := d.tPints[1:]                  // 未访问过的点
	visited := []string{d.tPints[0]}         // 访问过的点
	src := d.tPints[0]                       // 起点
	pre, next := src, src

	path := make(map[string][]string, 0)     // 起点到其它点的路径
	path[src + "_" + src] = []string{"A"}

	distanceGraph := make(map[string]int, 0)  // 起点到其它点的距离
	for len(tPoints) > 0 {
		distance := MaxDis
		var ind int = 0
		var dst string

		var nextInd int = 0

		for _, v := range visited {
			for ind, dst = range tPoints {
				newDis := d.tTwoPointDis[src + "_" + v] + d.tTwoPointDis[v + "_" + dst]  // 从起点src到已访问过的点v + 从v到未访问过点的距离
				if newDis < distance {
					distance = newDis
					pre = v
					next = dst
					nextInd = ind
					d.tTwoPointDis[src + "_" + dst] = distance
				}
			}
		}

		for _, tPoint := range path[src + "_" + pre] {
			path[src + "_" + next] = append(path[src + "_" + next], tPoint)
		}
		path[src + "_" + next] = append(path[src + "_" + next], next)  // 记录从src到next需经过的路径

		distanceGraph[src + "_" + next] = distance                     // 记录从src到next的距离

		visited = append(visited, next)
		tPoints = append(tPoints[:nextInd], tPoints[nextInd+1:]...)
	}

	fmt.Println(path)
	fmt.Println(distanceGraph)
}

func main() {
	d := new(Dijkstra)
	tPoints := []string{"A", "B", "C", "D"}
	tDis := [][]int{
		{0,      2, 6, 4},
		{127, 0, 3, 127},
		{7, 127, 0, 1},
		{5, 127, 12, 0}}

	d.Init(tPoints, tDis)
	d.dijkstra()
}
