package main

import "fmt"

/*
最大频率栈
*/
type FreqStack struct {
	stacks [][]int
	cnt    map[int]int
}

func Constructor() FreqStack {
	fs := FreqStack{cnt: map[int]int{}}

	return fs
}

func (this *FreqStack) Push(val int) {
	c := this.cnt[val]
	if c == len(this.stacks) {
		this.stacks = append(this.stacks, []int{val})
	} else {
		this.stacks[c] = append(this.stacks[c], val)
	}
	this.cnt[val] += 1
}

func (this *FreqStack) Pop() int {
	back := len(this.stacks) - 1
	tmpStk := this.stacks[back]
	res := tmpStk[len(tmpStk)-1]
	if len(tmpStk) == 1 {
		this.stacks = this.stacks[:back]
	} else {
		this.stacks[back] = tmpStk[:len(tmpStk)-1]
	}
	this.cnt[res] -= 1
	return res
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */

func main() {
	fs := Constructor()
	fs.Push(1)
	fs.Push(1)
	fs.Push(2)
	fmt.Println(fs.Pop())
	fmt.Println(fs.Pop())
	fmt.Println(fs.Pop())
}
