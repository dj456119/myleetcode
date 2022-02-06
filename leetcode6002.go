/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-06 10:56:42
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-06 12:01:25
 */
package myleetcode

type Bitset struct {
	BSet     []uint64
	OneCount int
	Size     int
}

func Constructor(size int) Bitset {
	z := size / 64
	z++
	m := []uint64{}
	for i := 0; i < z; i++ {
		m = append(m, 0)
	}
	return Bitset{m, 0, size}
}

func (this *Bitset) Fix(idx int) {
	pos := idx / 64
	pos1 := idx % 64
	z := this.BSet[pos]
	if z&(1<<pos1) == 0 {
		z = z | 1<<pos1
		this.BSet[pos] = z
		this.OneCount++
	}
}

func (this *Bitset) Unfix(idx int) {
	pos := idx / 64
	pos1 := idx % 64
	z := this.BSet[pos]
	if z&(1<<pos1) > 0 {
		z = ^z
		z = z | 1<<pos1
		z = ^z
		this.BSet[pos] = z
		this.OneCount--
	}
}

func (this *Bitset) Flip() {
	for i := range this.BSet {
		this.BSet[i] = ^this.BSet[i]
	}
	this.OneCount = this.Size - this.OneCount
}

func (this *Bitset) All() bool {
	return this.OneCount == this.Size
}

func (this *Bitset) One() bool {
	return this.OneCount > 0
}

func (this *Bitset) Count() int {
	return int(this.OneCount)
}

func (this *Bitset) ToString() string {
	result := []byte{}
	start := 64 - (this.Size % 64)
	for i := range this.BSet {

		m := this.BSet[i]

		f := make([]byte, 64)
		for z := 63; z >= 0; z-- {
			c := m % 2
			f[z] = byte(c) + '0'
			m = m / 2
		}
		result = append(f, result...)
	}
	m := result[start:]
	c1 := 0
	c2 := len(m) - 1
	for c1 < c2 {
		m[c1], m[c2] = m[c2], m[c1]
		c1++
		c2--
	}
	return string(m)
}
