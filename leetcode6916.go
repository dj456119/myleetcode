package main

func primeNumbers(n int) []int {
	if n < 2 {
		return []int{}
	}

	primes := []int{2}
	for i := 3; i <= n; i++ {
		isPrime := true
		sqrtI := int(math.Sqrt(float64(i)))

		for _, prime := range primes {
			if prime > sqrtI {
				break
			}
			if i%prime == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

func findPrimePairs(n int) [][]int {
	z := primeNumbers(n)
	result := [][]int{}
	i := 0
	j := len(z) - 1
	for {
		if i >= j {
			break
		}
		if z[i]+z[j] == n {
			result = append(result, []int{z[i], z[j]})
			continue
		}
		if z[i]+z[j] > n {
			j--
		} else {
			i++
		}
	}
}
