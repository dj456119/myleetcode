package myleetcode

func sumOfNumberAndReverse(num int) bool {
	for i := 0; i <= num; i++ {
		if i+int(resNum(i)) == num {
			return true
		}
	}
	return false
}

func resNum6219(n int) int64 {
	var result int64
	for n != 0 {
		z := n % 10
		result = result*10 + int64(z)
		n = n / 10
	}
	return result
}
