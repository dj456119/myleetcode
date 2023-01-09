package myleetcode

func countPrimes(n int) int {
	isPrimes := make([]bool, n)
	for i := range isPrimes {
		isPrimes[i] = true
	}
	result := 0
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			result++
			for j := 2; j*i < n; j++ {
				isPrimes[j*i] = false
			}
		}
	}
	return result
}
