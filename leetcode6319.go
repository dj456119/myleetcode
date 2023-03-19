package myleetcode

func evenOddBit(n int) []int {
	z := []int{}
	for n != 0 {
		z = append(z, n%2)
		n = n / 2
	}
	jishu := 0
	oushu := 0
	for i := range z {
		if i%2 == 0 && z[i] == 1 {
			oushu++
		}
		if i%2 == 1 && z[i] == 1 {
			jishu++
		}
	}
	return []int{oushu, jishu}
}
