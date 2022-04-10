package myleetcode

func minimizeResult(expression string) string {

}

func backtrace(str []byte, pos int, isLeft, isRight bool, isadd bool, first, second, third, fourth int, result []byte, min *int, ss *string) {
	if pos == len(str) {
		if isLeft && !isRight {
			if fourth == -1 {
				third = 1
			}
			if first == -1 {
				first = 1
			}
			z := first * (second + third) * fourth
			if z < *min {
				*min = z
				*ss = string(append(result, ')'))
			}
			return
		}
		if !isLeft {
			return
		}
		if isLeft && isRight {
			if fourth == -1 {
				third = 1
			}
			if first == -1 {
				first = 1
			}
			z := first * (second + third) * fourth
			if z < *min {
				*min = z
				*ss = string(append(result, ')'))
			}
		}
	}
	//	if str[pos] == '+'
}
