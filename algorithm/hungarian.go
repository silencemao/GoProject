package main

import "fmt"

/*
匈牙利算法 - 先到先得 能让则让
给定两个点集 [A B C D] [a b c d]
一个邻接矩阵matrix，矩阵matrix[i][j] >=0 表示点i-j之间有邻接关系，<0 表示没有关系
根据邻接矩阵的关系，返回最大匹配值，即有几条边可以连接两个集合中的点，每个点只归属于一条边
https://blog.csdn.net/dark_scope/article/details/8880547
https://zhuanlan.zhihu.com/p/62981901
*/

func match_hungarian(i, n int, vis []bool, p []int, matrix [][]int) bool {
	for j := 0; j < n; j++ {
		if matrix[i][j] >= 0 && !vis[j] { // 有关系 并且没有访问过
			vis[j] = true
			if p[j] == -1 || match_hungarian(p[j], n, vis, p, matrix) { // 右边点未分配 或者 为分配的点重新尝试找新的位置
				p[j] = i
				return true
			}
		}
	}
	return false
}

func hungarian() {
	m, n := 4, 4       // 两个集合的点数
	matrix := [][]int{ // 邻接关系
		{1, -1, 1, -1},
		{-1, 1, -1, -1},
		{1, -1, 1, 1},
		{1, 1, -1, -1},
	}
	vis := make([]bool, n) // 是否访问过左边点
	p := make([]int, n)    // 右边集合对应左边集合的点

	for i := range p { // 先标记没有任何对应关系
		p[i] = -1
	}

	cnt := 0
	for i := 0; i < m; i++ {
		vis = make([]bool, n)
		if match_hungarian(i, n, vis, p, matrix) {
			cnt++
		}
	}
	fmt.Println(p)
	fmt.Println(cnt)
}

/*
https://www.cnblogs.com/zpfbuaa/p/7218607.html

https://www.cnblogs.com/DarrenChan/p/10568268.html
*/

func main() {
	hungarian()
	km_main()
}

// 下面是KM算法
const MAX = 100
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

//const INT_MAX = 1061109567
//const INT_MIN = -1061109567
var sx []bool //记录寻找增广路时点集x，y里的点是否搜索过
var sy []bool
var match []int //match[i]记录y[i]与x[match[i]]相对应
var weight [][]int
var Cx []int
var Cy []int

func search_path(u int) bool { //给x[u]找匹配,这个过程和匈牙利匹配是一样的
	sx[u] = true
	for v := 0; v < len(weight[0]); v++ {
		if !sy[v] && Cx[u]+Cy[v] == weight[u][v] {
			sy[v] = true
			if match[v] == -1 || search_path(match[v]) { //如果第v个y点还没被占，或者第v个y点还可以找到其他可搭配的x点
				match[v] = u
				return true
			}
		}
	}
	return false
}

// weight是权重
func Kuhn_Munkras(max_weight bool) int {
	//如果求最小匹配，则要将边权取反
	if !max_weight {
		for i := 0; i < len(weight); i++ {
			for j := 0; j < len(weight[0]); j++ {
				weight[i][j] = -weight[i][j]
			}
		}
	}

	//初始化顶标，Cx[i]设置为max(weight[i][j] | j=0,..,n-1 ), Cy[i]=0;
	//Cy的顶标都是0
	for i := 0; i < len(Cx); i++ {
		Cx[i] = INT_MIN
		for j := 0; j < len(weight[0]); j++ {
			if Cx[i] < weight[i][j] {
				Cx[i] = weight[i][j]
			}
		}
	}
	fmt.Println(Cx, Cy)
	for i := 0; i < len(match); i++ {
		match[i] = -1
	}

	//不断修改顶标，直到找到完备匹配或完美匹配
	for u := 0; u < len(weight); u++ { //为x里的每一个点找匹配
		for {
			for index, _ := range sx {
				sx[index] = false
			}
			for index, _ := range sy {
				sy[index] = false
			}
			if search_path(u) { //x[u]在相等子图找到了匹配,继续为下一个点找匹配
				break
			}

			//如果在相等子图里没有找到匹配，就修改顶标，直到找到匹配为止
			//首先找到修改顶标时的增量inc, min(Cx[i] + Cy [i] - weight[i][j],inc);,Cx[i]为搜索过的点，Cy[i]是未搜索过的点,因为现在是要给u找匹配，所以只需要修改找的过程中搜索过的点，增加有可能对u有帮助的边
			inc := INT_MAX
			for i := 0; i < len(weight); i++ {
				if sx[i] {
					for j := 0; j < len(weight[0]); j++ {
						if !sy[j] && (Cx[i]+Cy[j]-weight[i][j]) < inc {
							inc = Cx[i] + Cy[j] - weight[i][j]
						}
					}
				}
			}

			//找不到可以加入的边，返回失败（即找不到完美匹配）
			if inc == INT_MAX {
				return -1
			}

			//找到增量后修改顶标，因为sx[i]与sy[j]都为真，则必然符合Cx[i] + Cy [j] =weight[i][j]，然后将Cx[i]减inc，Cy[j]加inc不会改变等式，但是原来Cx[i] + Cy [j] ！=weight[i][j]即sx[i]与sy[j]最多一个为真，Cx[i] + Cy [j] 就会发生改变，从而符合等式，边也就加入到相等子图中
			for i := 0; i < len(weight); i++ {
				if sx[i] { //如果点x在S集合里
					Cx[i] -= inc
				}
			}
			for j := 0; j < len(weight[0]); j++ {
				if sy[j] { //如果点y在T集合里
					Cy[j] += inc
				}
			}
		}

	}
	sum := 0
	for i := 0; i < len(weight[0]); i++ {
		if match[i] > -1 {
			sum += weight[match[i]][i]
		}
	}

	if !max_weight {
		sum = -sum
		return sum
	}
	return sum
}

func km_main() {
	sx = make([]bool, 3)
	sy = make([]bool, 3)
	match = make([]int, 3)
	weight = [][]int{
		{3, 0, 4},
		{2, 1, 3},
		{0, 0, 5},
	}

	Cx = make([]int, 3)
	Cy = make([]int, 3)

	sum := Kuhn_Munkras(true)
	fmt.Println(sum)

	fmt.Println(match)
}
