package jianzhioffer

// 用两个栈实现队列
type CQueue struct {
	stackIn, stackOut []int
}

func CQueueConstructor() CQueue {
	return CQueue{stackIn: []int{}, stackOut: []int{}}
}

func (this *CQueue) AppendTail(value int) {
	this.stackIn = append(this.stackIn, value)
}

func (this *CQueue) DeleteHead() int {
	res := int(-1)
	if len(this.stackOut) == 0 {
		for len(this.stackIn) > 0 {
			temp := this.stackIn[len(this.stackIn)-1]
			this.stackOut = append(this.stackOut, temp)
			this.stackIn = this.stackIn[:len(this.stackIn)-1]
		}
	}
	if len(this.stackOut) > 0 {
		res = this.stackOut[len(this.stackOut)-1]
		this.stackOut = this.stackOut[:len(this.stackOut)-1]
	}
	return res
}
