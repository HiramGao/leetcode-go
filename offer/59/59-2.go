package main

type Data struct {
	index int
	value int
}
type MaxQueue struct {
	Q, MaxQ []Data
	Current int
}

func Constructor() MaxQueue {
	return MaxQueue{
		Q:       []Data{},
		MaxQ:    []Data{},
		Current: 0,
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.Q) == 0 {
		return -1
	}
	return this.MaxQ[0].value
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.MaxQ) != 0 && value >= this.MaxQ[len(this.MaxQ)-1].value {
		this.MaxQ = this.MaxQ[:len(this.MaxQ)-1]
	}
	tmp := Data{value: value, index: this.Current}
	this.Q = append(this.Q, tmp)
	this.MaxQ = append(this.MaxQ, tmp)
	this.Current++
}

func (this *MaxQueue) Pop_front() int {
	if len(this.Q) == 0 {
		return -1
	}
	if this.MaxQ[0].index == this.Q[0].index {
		this.MaxQ = this.MaxQ[1:]
	}
	tmp := this.Q[0]
	this.Q = this.Q[1:]
	return tmp.value
}

func main() {

}
