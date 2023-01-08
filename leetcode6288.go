package myleetcode

type DataStream struct {
	St    []int
	Value int
	K     int
	Lx    int
}

func Constructor6288(value int, k int) DataStream {
	ds := DataStream{St: []int{}, Value: value, K: k, Lx: 0}
	return ds
}

func (this *DataStream) Consec(num int) bool {
	if num == this.Value {
		this.Lx = this.Lx + 1
	} else {
		this.Lx = 0
	}
	if this.Lx >= this.K {
		return true
	}
	return false
}

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */
