package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
)

func main() {
	a := make(map[string][]int)

	b := []int{1}
	if _, ok := a["a"]; !ok {
		a["a"] = b
	}

	fmt.Println(a)

	calculateAngle()
	//isZero()

	//testRand()
	//
	//testRand1()
	//
	//testListClear()


}
// 与Y轴夹角
func calculateAngle() {

	a := []int{1, 1}
	b := []int{2, -4}
	tmpAngle := math.Atan2(float64(b[0] - a[0]), float64(b[1] - a[1]))
	fmt.Println(tmpAngle)
	tmpAngle = 180 * tmpAngle / math.Pi
	fmt.Println(tmpAngle)
}

func isZero() {
	a := 0.000000
	fmt.Println(a==0.0)
}

func testRand() {
	a := rand.Float64()
	p := 2.0
	res := math.Pow(a, p)
	fmt.Println(a, res)
}

func testRand1() {
	r := rand.New(rand.NewSource(0))
	fmt.Println(reflect.TypeOf(r))

}

func testListClear() {
	var a []int
	for i := 0; i < 5; i++ {
		a =append(a, i)
	}
	fmt.Println(a, len(a), cap(a))
	a = nil
	fmt.Println(a, len(a), cap(a))

	a = append(a, 10)
	a = append(a, 15)
	fmt.Println(a, len(a), cap(a))
}
