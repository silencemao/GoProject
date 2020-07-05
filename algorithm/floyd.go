package main

import (
	"fmt"
	"strconv"
)

//https://juejin.im/post/5cc79c93f265da035b61a42e

type Floyd struct {
	tTwoPointDis [][]int
	tPath [][]int
}

func (f *Floyd) Init(tDis [][]int) {
	f.tTwoPointDis = tDis

	r := len(tDis)

	f.tPath = make([][]int, r)
	for i := range f.tPath {
		f.tPath[i] = make([]int, r)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			f.tPath[i][j] = -1
		}
	}
}

func (f *Floyd) solve() {
	fmt.Println("before")
	for _, tNums := range f.tTwoPointDis {
		for _, tNum := range tNums {
			fmt.Print(tNum, " ")
		}
		fmt.Println()
	}
	r := len(f.tTwoPointDis)
	for k := 0; k < r; k++ {
		for i := 0; i < r; i++ {
			for j := 0; j < r; j++ {
				if f.tTwoPointDis[i][j] > (f.tTwoPointDis[i][k] + f.tTwoPointDis[k][j]) {
					f.tPath[i][j] = k
					f.tTwoPointDis[i][j] = f.tTwoPointDis[i][k] + f.tTwoPointDis[k][j]
				}
			}
		}
	}
	fmt.Println("after")
	for _, tNums := range f.tTwoPointDis {
		for _, tNum := range tNums {
			fmt.Print(tNum, " ")
		}
		fmt.Println()
	}

	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			if i != j {
				fmt.Println(f.getPath(i, j))
			}
		}
	}
}

func (f *Floyd) getPath(i, j int) string {
	if f.tPath[i][j] == -1 {
		return " " + strconv.Itoa(i) + " " + strconv.Itoa(j)
	} else {
		k := f.tPath[i][j]
		return f.getPath(i, k) + f.getPath(k, j)
	}
}

func main() {
	tDis := [][]int{
		{0,      2, 6, 4},
		{127, 0, 3, 127},
		{7, 127, 0, 1},
		{5, 127, 12, 0}}

	f := new(Floyd)
	f.Init(tDis)
	f.solve()
}


