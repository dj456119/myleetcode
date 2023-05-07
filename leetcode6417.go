package myleetcode

type FrequencyTracker struct {
	Number []int
	Count  []map[int]bool
}

func Constructor() FrequencyTracker {
	f := FrequencyTracker{Number: make([]int, 100001), Count: make([]map[int]bool, 100001)}
	return f
}

func (this *FrequencyTracker) Add(number int) {
	this.Number[number]++
	z := this.Number[number]
	if this.Count[z] == nil {
		this.Count[z] = make(map[int]bool)
	}
	this.Count[z][number] = true
	delete(this.Count[z-1], number)
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.Number[number] == 0 {
		return
	}
	z := this.Number[number]
	if z == 1 {
		delete(this.Count[1], number)
		this.Number[number]--
		return
	}
	delete(this.Count[z], number)
	this.Count[z-1][number] = true
	this.Number[number]--
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return len(this.Count[frequency]) > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
