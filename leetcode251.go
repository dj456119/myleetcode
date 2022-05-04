package myleetcode

type Vector2D struct {
	vec [][]int
	Y   int
	X   int
}

func Constructor251(vec [][]int) Vector2D {
	v2 := Vector2D{vec: vec}
	return v2
}

func (this *Vector2D) Next() int {
	// fmt.Println(this)
	for {
		if this.Y == len(this.vec) {
			return -1
		}
		if len(this.vec[this.Y]) == 0 {
			this.Y = this.Y + 1
		} else {
			break
		}
	}
	c := this.vec[this.Y][this.X]
	if this.X+1 == len(this.vec[this.Y]) {
		this.Y = this.Y + 1
		this.X = 0
	} else {
		this.X = this.X + 1
	}
	return c
}

func (this *Vector2D) HasNext() bool {
	for {
		if this.Y == len(this.vec) {
			return false
		}
		if len(this.vec[this.Y]) == 0 {
			this.Y = this.Y + 1
		} else {
			break
		}
	}
	if this.Y == len(this.vec) {
		return false
	}
	if len(this.vec[this.Y]) == 0 {
		return false
	}
	return true
}
