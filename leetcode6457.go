package myleetcode

func removeTrailingZeros(num string) string {
	for len(num) > 0 && num[len(num)-1] == '0' {
		num = num[:len(num)-1]
	}
	return num
}
