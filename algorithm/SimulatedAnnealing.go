package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	x := rand.Float64() * 100
	bestX := x
	fmt.Println(bestX)

	for t > minT && iterNum >= 0 {
		funcTmp := getFuncRes(x, 0)
		xNew := x + rand.Float64() * 2 - 1
		if xNew >= 0 && xNew <= 100 {
			funcNew := getFuncRes(xNew, 0)
			if funcNew < funcTmp {
				x = xNew
				bestX = x
				result = funcNew
			} else {
				p := math.Exp(-1 * (funcNew - funcTmp) / t)
				if rand.Float64() < p {
					x = xNew
				}
				if p > 0 {
					fmt.Println("prob ", p)
				}

			}
		}
		iterNum--
		t = t * delta
	}
	fmt.Println(bestX, result)
}

func main() {
	SimulateAnneal()
}
