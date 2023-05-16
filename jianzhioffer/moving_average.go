package jianzhioffer

// 滑动窗口的平均值
//给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算滑动窗口里所有数字的平均值。
//
//实现 MovingAverage 类：
//
//MovingAverage(int size) 用窗口大小 size 初始化对象。
//double next(int val) 成员函数 next 每次调用的时候都会往滑动窗口增加一个整数，请计算并返回数据流中最后 size 个值的移动平均值，即滑动窗口里所有数字的平均值。

type MovingAverage struct {
	sum   int
	size  int
	queue []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.queue) == this.size && len(this.queue) > 0 {
		// queue已经满了，需要移除队列头部元素
		this.sum -= this.queue[0]
		this.queue = this.queue[1:len(this.queue)]
	}
	this.sum += val
	this.queue = append(this.queue, val)
	return float64(this.sum) / float64(len(this.queue))
}
