package myleetcode

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	z := 0
	return backtrace6029(numArrows, aliceArrows, 0, 0, &z, []int{})
}

func backtrace6029(nnumArrows int, aliceArrows []int, pos int, score int, max *int, result []int) []int {
	if pos == len(aliceArrows) {
		if score > *max {
			*max = score
			return result
		}
		return nil
	}

	result1 := make([]int, len(result))
	copy(result1, result)
	r := backtrace6029(nnumArrows, aliceArrows, pos+1, score, max, result)

	if aliceArrows[pos] < nnumArrows {
		r = backtrace6029(nnumArrows-aliceArrows[pos], aliceArrows, pos+1, score+pos, max, append(result1, aliceArrows[pos]))
	}
	return r
}
