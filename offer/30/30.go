package main

import "math"

type MinStack struct {
	stack    []int
	MinStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		MinStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.MinStack[len(this.MinStack)-1]
	this.MinStack = append(this.MinStack, min(top, x))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.MinStack = this.MinStack[:len(this.MinStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.MinStack[len(this.MinStack)-1]
}

func min(y int, x int) int {
	if y < x {
		return y
	}
	return x
}

func main() {

}
