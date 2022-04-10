package myleetcode

func largestInteger(num int) int {
	max := -1
	z := []int{}
	c := num
	for c != 0 {
		z = append(z, c%10)
		c = c / 10
	}
	start := 0
	end := len(z) - 1
	for start < end {
		z[start], z[end] = z[end], z[start]
		start++
		end--
	}
	for i := range z {
		if z[i]%2 == 0 && z[i] != 0 && i != len(z)-1 {
			for j := i + 1; j < len(z); j++ {
				if z[j]%2 == 0 && z[i] != 0 {
					z[i], z[j] = z[j], z[i]
					f := createNum(z)
					if f > max {
						max = f
					}
					z[j], z[i] = z[i], z[j]
				}
			}
		}
		if z[i]%2 != 0 && i != len(z)-1 {
			for j := i + 1; j < len(z); j++ {
				if z[j]%2 != 0 {
					z[i], z[j] = z[j], z[i]
					f := createNum(z)
					if f > max {
						max = f
					}
					z[j], z[i] = z[i], z[j]
				}
			}
		}
	}
	return max
}

func createNum(c []int) int {
	m := 0
	for i := range c {
		m = m*10 + c[i]
	}
	return m
}
