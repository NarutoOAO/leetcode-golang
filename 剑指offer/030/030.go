package _30

import "math"

type MinStack struct {
	stack1 []int
	stack2 []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	s1, s2 := make([]int, 0), make([]int, 0)
	s2 = append(s2, math.MaxInt64)
	return MinStack{stack1: s1, stack2: s2}
}

func (this *MinStack) Push(x int) {
	this.stack1 = append(this.stack1, x)
	if x <= this.stack2[len(this.stack2)-1] {
		this.stack2 = append(this.stack2, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack1) > 0 {
		if this.stack2[len(this.stack2)-1] == this.stack1[len(this.stack1)-1] {
			this.stack2 = this.stack2[:len(this.stack2)-1]
		}
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack1[len(this.stack1)-1]
}

func (this *MinStack) Min() int {
	if len(this.stack2) == 1 {
		return 0
	}
	return this.stack2[len(this.stack2)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
