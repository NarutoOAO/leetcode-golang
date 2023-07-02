package _59

// 使用两个队列，每次加入新元素的时候将队列前面比新元素小的元素弹出
type MaxQueue struct {
	q1 []int
	q2 []int
}

func Constructor() MaxQueue {
	return MaxQueue{q1: make([]int, 0), q2: make([]int, 0)}
}

func (this *MaxQueue) Max_value() int {
	if len(this.q2) == 0 {
		return -1
	}
	return this.q2[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q1 = append(this.q1, value)
	for len(this.q2) > 0 && value > this.q2[len(this.q2)-1] {
		this.q2 = this.q2[:len(this.q2)-1]
	}
	this.q2 = append(this.q2, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q1) == 0 {
		return -1
	}
	v := this.q1[0]
	if v == this.q2[0] {
		this.q2 = this.q2[1:len(this.q2)]
	}
	this.q1 = this.q1[1:len(this.q1)]
	return v
}
