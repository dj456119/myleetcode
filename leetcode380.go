package myleetcode

import "math/rand"

type RandomizedSet struct {
	Data      map[int]int
	DataArray []int
}

func Constructor380() RandomizedSet {
	rs := RandomizedSet{Data: make(map[int]int), DataArray: []int{}}
	return rs
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Data[val]; ok {
		return false
	}
	this.Data[val] = len(this.DataArray)
	this.DataArray = append(this.DataArray, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if pos, ok := this.Data[val]; !ok {
		return false
	} else {
		delete(this.Data, val)
		this.DataArray[len(this.DataArray)-1], this.DataArray[pos] = this.DataArray[pos], this.DataArray[len(this.DataArray)-1]
		this.DataArray = this.DataArray[:len(this.DataArray)-1]
		if pos == len(this.DataArray) {
			return true
		}
		this.Data[this.DataArray[pos]] = pos
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.DataArray[rand.Intn(len(this.DataArray))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
