/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 16:07:04
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 16:27:32
 */
package myleetcode

type MinStack struct {
	stack    []int
	min      int
	init_min bool
}

func Constructor155() MinStack {
	return MinStack{stack: []int{}, init_min: true}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.init_min || this.min > val {
		this.min = val
		this.init_min = false
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) != 0 {
		nMin := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		if len(this.stack) == 0 {
			this.min = 0
			this.init_min = true
		} else {
			if nMin == this.min {
				this.init_min = true
				for _, j := range this.stack {
					if this.init_min || j < this.min {
						this.min = j
						this.init_min = false
					}
				}
			}
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) != 0 {
		return this.stack[len(this.stack)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
