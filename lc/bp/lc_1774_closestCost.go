package main

import (
	"fmt"
	"math"
)

/*
你打算做甜点，现在需要购买配料。目前共有 n 种冰激凌基料和 m 种配料可供选购。而制作甜点需要遵循以下几条规则：

必须选择 一种 冰激凌基料。
可以添加 一种或多种 配料，也可以不添加任何配料。
每种类型的配料 最多两份 。
给你以下三个输入：
baseCosts ，一个长度为 n 的整数数组，其中每个 baseCosts[i] 表示第 i 种冰激凌基料的价格。
toppingCosts，一个长度为 m 的整数数组，其中每个 toppingCosts[i] 表示 一份 第 i 种冰激凌配料的价格。
target ，一个整数，表示你制作甜点的目标价格。
你希望自己做的甜点总成本尽可能接近目标价格 target 。

返回最接近 target 的甜点成本。如果有多种方案，返回成本相对较低 的一种。

给定两个数组 basecost、toppingcost以及目标值target
basecost中必须选一个值，topping中可以选也可不选，topping中的每个值最多选两次，返回和最接近target的结果
*/
func help1774(toppingCosts []int, tmp []int, target, ind, curSum int, res *int) {
	if ind >= len(toppingCosts) {
		return
	}
	if math.Abs(float64(curSum-target)) < math.Abs(float64(*res-target)) {
		*res = curSum
	} else if math.Abs(float64(curSum-target)) == math.Abs(float64(*res-target)) {
		if curSum < *res {
			*res = curSum
		}
	}

	for i := ind; i < len(toppingCosts); i++ {
		if tmp[i] > 1 || math.Abs(float64(curSum+toppingCosts[i]-target)) > math.Abs(float64(curSum-target)) {
			continue
		}
		tmp[i] += 1
		curSum += toppingCosts[i]
		help1774(toppingCosts, tmp, target, i, curSum, res)
		tmp[i] -= 1
		curSum -= toppingCosts[i]
	}
}

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	final, diff := 0, 1<<31-1

	for _, b := range baseCosts {
		res := 1 << 31
		curSum := b
		tmp := make([]int, len(toppingCosts))
		help1774(toppingCosts, tmp, target, 0, curSum, &res)
		fmt.Println(res, final, diff)
		if int(math.Abs(float64(res-target))) < diff {
			final = res
			diff = int(math.Abs(float64(res - target)))
		} else if int(math.Abs(float64(res-target))) == diff && res < final {
			final = res
		}
	}
	return final
}

func main() {
	baseCosts := []int{3, 10}
	toppingCosts := []int{2, 5}
	target := 9
	fmt.Println(closestCost(baseCosts, toppingCosts, target))
}
