package myleetcode

func myPow(x float64, n int) float64 {
	var result float64 = 1
	var q float64 = x
	a := n
	for n != 0 {
		z := n % 2
		if z == 0 {
			q = q * q
		} else {
			result = result * q
			q = q * q
		}
		n = n / 2
	}
	if a < 0 {
		return 1 / result
	}
	return result
}
