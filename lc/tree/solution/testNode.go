package main

import "fmt"

type Node struct {
	Val int
	Next *Node
}

// https://www.guaosi.com/2019/03/20/golang-value-passing-and-reference-passing/

func main() {
	pHead := &Node{
		Val:  10,
		Next: nil,
	}
	fmt.Printf("%p\n", pHead)    // pHead指向的地址
	fmt.Printf("%p\n", &pHead)   // pHead本身所在的位置

	var pNodes []*Node
	for i := 0; i < 5; i++ {
		pNode := &Node{
			Val:  i,
		}
		pNodes = append(pNodes, pNode)
	}

	pHead.Next = pNodes[0]

	for i := 0; i < len(pNodes)-1; i++ {
		pNodes[i].Next = pNodes[i+1]
	}

	//m := test(pHead)
	//fmt.Println(m.Val)
	fmt.Println(pHead.Val)
	test(pHead)
	fmt.Println(pHead.Val)
	fmt.Println(pHead.Next.Next.Val)
	fmt.Printf("%p\n", pHead)    // pHead所指向的地址没有变

	//transporter := make(map[int]bool)
	//for i := 0; i < 5; i++ {
	//	test1(i, transporter)
	//}
	//
	//fmt.Println(transporter)
}

func test(n *Node) {
	fmt.Printf("%p\n", n)
	fmt.Printf("%p\n", &n)
	n = n.Next
	n = n.Next
	n.Val = 100
	fmt.Printf("%p\n", n)
	//return n
}

func test1(i int, trans map[int]bool) {
	trans[i-1] = false
	trans[i] = true
}
