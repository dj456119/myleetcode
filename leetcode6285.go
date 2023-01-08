package myleetcode

import (
	"container/heap"
	"math"
	"sort"
)

func maxKelements(nums []int, k int) int64 {
	h := &Hp{nums}
	heap.Init(h)
	var result int64
	for i := 0; i < k; i++ {
		z := heap.Pop(h).(int)
		var d float64
		d = float64(z) / float64(3)
		z = int(math.Ceil(d))
		result = result + int64(z)
		heap.Push(h, z)
	}
	return result
}

type Hp struct {
	H sort.IntSlice
}

func (hp Hp) Less(i, j int) bool {
	return hp.H[i] > hp.H[j]
}

func (hp Hp) Len() int {
	return len(hp.H)
}

func (hp Hp) Swap(i, j int) {
	hp.H[i], hp.H[j] = hp.H[j], hp.H[i]
}

func (hp *Hp) Push(x interface{}) {
	hp.H = append(hp.H, x.(int))
}

func (hp *Hp) Pop() interface{} {
	z := hp.H[len(hp.H)-1]
	hp.H = hp.H[:len(hp.H)-1]
	return z
}
