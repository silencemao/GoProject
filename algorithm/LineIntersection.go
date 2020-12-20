package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func getParam(p1 Point, p2 Point) (float64, float64, float64) {
	a := p2.y - p1.y
	b := p1.x - p2.x
	c := p2.x * p1.y - p1.x * p2.y
	return a, b, c
}

func isParallel(a1, b1, a2, b2 float64) bool {
	if b1 * a2 == b2 * a1 {
		return true
	}
	return false
}

/*
首先设交点坐标为 (x, y)，两线段对应直线的一般式为：
a1x + b1y + c1 = 0
a2x + b2y + c2 = 0
那么对 1 式乘 a2，对 2 式乘 a1 得：
a2*a1x + a2*b1y + a2*c1 = 0
a1*a2x + a1*b2y + a1*c2 = 0
两式相减得：
y = (c1 * a2 - c2 * a1) / (a1 * b2 - a2 * b1)

同样可以推得：
x = (c2 * b1 - c1 * b2) / (a1 * b2 - a2 * b1)
 */

func getIntersection(a1, b1, c1, a2, b2, c2 float64) Point {
	x := (c2 * b1 - c1 * b2) / (a1 * b2 - a2 * b1)
	y := (c1 * a2 - c2 * a1) / (a1 * b2 - a2 * b1)
	p := Point{x, y}
	return p
}

func isPointInLine(pIntersection, pStartPoint, pEndPoint Point) bool {
	d1 := math.Sqrt(math.Pow(float64(pIntersection.x-pStartPoint.x), 2) + math.Pow(float64(pIntersection.y-pStartPoint.y), 2))
	d2 := math.Sqrt(math.Pow(float64(pIntersection.x-pEndPoint.x), 2) + math.Pow(float64(pIntersection.y-pEndPoint.y), 2))
	d3 := math.Sqrt(math.Pow(float64(pStartPoint.x-pEndPoint.x), 2) + math.Pow(float64(pStartPoint.y-pEndPoint.y), 2))
	if math.Abs(d1 + d2 - d3) <= 0.0001 {
		return true
	}
	return false
}

func main() {
	pStart1 := Point{1, 5}
	pEnd1 := Point{4, 0}

	pStart2 := Point{2, 0}
	pEnd2 := Point{5, 3}

	a1, b1, c1 := getParam(pStart1, pEnd1)
	a2, b2, c2 := getParam(pStart2, pEnd2)

	if isParallel(a1, b1, a2, b2) {
		fmt.Println("两条线段平行")
	}

	pIntersection := getIntersection(a1, b1, c1, a2, b2, c2)

	if isPointInLine(pIntersection, pStart1, pEnd1) && isPointInLine(pIntersection, pStart2, pEnd2) {
		fmt.Println("交点为：", pIntersection.x, pIntersection.y)
	} else {
		fmt.Println("无交点")
	}
}
