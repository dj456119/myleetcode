package myleetcode

type NumMatrix struct {
	Mmap   [][]int
	Matrix [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
	Nm := NumMatrix{}
	Nm.Matrix = matrix
	mmap := make([][]int, len(matrix)+1)
	for i := range mmap {
		mmap[i] = make([]int, len(matrix[0])+1)
	}
	for i := range matrix {
		rx := 0
		for j := range matrix[i] {
			rx = matrix[i][j] + rx
			mmap[i+1][j+1] = rx + mmap[i][j+1]
		}
	}
	Nm.Mmap = mmap
	return Nm
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.Mmap[row2+1][col2+1] + this.Mmap[row1][col1] - this.Mmap[row2+1][col1] - this.Mmap[row1][col2+1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
