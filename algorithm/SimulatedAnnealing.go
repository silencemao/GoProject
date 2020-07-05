package main

import (
	"fmt"
	"math"
	"math/rand"
	"myGoProject/algorithm/heuristic"
)

func getFuncRes(x, y float64) float64{
	return 6.0 * math.Pow(x, 7) + 8.0 * math.Pow(x, 6)  + 7.0 * math.Pow(x, 3) + 5 * math.Pow(x, 2) - x*y
}

func SimulateAnneal() {
	result := math.MaxFloat64
	t := 100.0
	minT := 1e-8
	iterNum := 10000
	delta := 0.98
	rand.Seed(0)
	x := rand.Float64() * 100
	bestX := x
	fmt.Println(bestX)

	cnt := 0
	for t > minT && iterNum >= 0 {
		xNew := x + rand.Float64() * 2 - 1
		if xNew >= 0 && xNew <= 100 {
			cnt++
			funcNew := getFuncRes(xNew, 0)
			if funcNew < result {
				x = xNew
				bestX = x
				result = funcNew
			} else {
				p := math.Exp(-1 * (funcNew - result) / t)
				if rand.Float64() < p {
					x = xNew
				}
				if p > 0 {
					//fmt.Println("prob ", p)
				}

			}
		}
		iterNum--
		t = t * delta
	}
	fmt.Println(bestX, result, cnt)
}

func LateAcc() {

	rand.Seed(0)

	bestX := rand.Float64() * 100
	bestRes := 0.0

	pLa := new(heuristic.LateAcceptance)
	pLa.Init(200)

	pIterNum := 10000
	cnt := 0
	for pIterNum > 0 {
		x := bestX + rand.Float64() * 2 - 1
		if x >= 0 && x <= 100 {
			cnt++
			pRes := getFuncRes(x, 0)
			if pLa.Accept(pRes) {
				bestX = x
				bestRes = pRes
			}
		}
		pIterNum--
	}
	fmt.Println(bestX, bestRes, cnt)
}

func main() {
	SimulateAnneal()
	LateAcc()
}
