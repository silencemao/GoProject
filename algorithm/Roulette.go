package main

import (
	"fmt"
	"math/rand"
	"time"
)
// 轮盘赌算法
// 随机生成n个适应度值
func getAdaptability(n int) []float32 {
	rand.Seed(time.Now().UnixNano())
	var syd []float32
	for i :=0 ; i < n; i++ {
		syd = append(syd, rand.Float32())
	}
	return syd
}

/*
1、首先对适应度值求和，同时生成一个长度为len(syd)的计数数组
2、对每一个适应度值求概率，并对概率进行累加，存入sydPro中。sydPro中的每一个概率值是前面所有概率值的和，最后一个值为1
*/
func Roulette() {
	syd := getAdaptability(10)
	var sydSum float32
	var countPerSyd []int
	for i :=0 ; i < len(syd); i++ {
		sydSum += syd[i]
		countPerSyd = append(countPerSyd, 0)
	}
	var sydPro []float32
	sydPro = append(sydPro, syd[0] / sydSum)
	for i := 1; i < len(syd); i++ {
		sydPro = append(sydPro, sydPro[i-1] + syd[i] / sydSum)
	}

	iterNum := 10000
	for iterNum > 0 {
		p := rand.Float32()
		for i := 0; i < len(sydPro); i++ {
			if p <= sydPro[i] {
				countPerSyd[i]++
				break
			}
		}
		iterNum--
	}

	fmt.Println(syd)
	fmt.Println(sydPro)
	fmt.Println(countPerSyd)
}

func main() {
	Roulette()
}
