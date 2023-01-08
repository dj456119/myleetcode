package myleetcode

func smallestValue(n int) int {
	for {
		n1 := fenjie(n)
		if n == n1 {
			return n
		}
		n = n1
	}
	return -1
}

func fenjie(n int) int {
	result := 0
	for {
		if n == 1 {
			return result
		}
		for i := 2; i <= n; i++ {
			if n%i == 0 {
				result = result + i
				n = n / i
				break
			}
		}
	}

}
