package main

import (
	"fmt"
	"strconv"
)

/*
-1 right
-2 left
目标最初在坐标(0, 0)处
给定一个数组commands表示获得命令，命令可以为-1，-2或1-9的数字，-1表示向右转90度，-2表示向左转90度，1-9表示在当前方向上前进的步数
obstacles表示障碍物的坐标，当前进的方向上有障碍物的时候，会停在障碍物的前面，
返回目标可以走的最远距离，即距离坐标(0, 0)的最远距离

dirs的方向设计的比较巧妙，最开始目标面向北(d=0)，即y+方向，
                      向右转即改为x+方向(d=1)，dir[1][0]=1 dir[1][1]=0即x会增加，y不变，则沿着x+方向前进了

*/
func robotSim(commands []int, obstacles [][]int) int {
	res := 0
	set := make(map[string]int)
	dirs :=[][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}  // 方向，即区分x+、x-、y+、y-
	d, x, y := 0, 0, 0
	for _, obstacle := range obstacles {  // 标记障碍物的位置
		set[strconv.Itoa(obstacle[0]) + "_" + strconv.Itoa(obstacle[1])] = 1
	}

	for _, command :=range commands {
		if command == -1 {
			d++
			if d==4 {
				d = 0
			}
		} else if command == -2 {
			d--
			if d==-1 {
				d = 3
			}
		} else {
			for command > 0 {
				command--
				_, ok := set[strconv.Itoa(x + dirs[d][0]) + "_" + strconv.Itoa(y + dirs[d][1])]
				if !ok {
					x += dirs[d][0]
					y += dirs[d][1]
				}
			}
			if x * x + y * y > res {
				res = x * x + y * y
			}
		}
	}
	return res
}

func main() {
	commands := []int{4, -1, 4, -2, 4}
	obstacles := [][]int{{2, 4}}
	fmt.Println(robotSim(commands, obstacles))
}