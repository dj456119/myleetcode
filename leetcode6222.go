package myleetcode

func makeIntegerBeautiful(n int64, target int) int64 {
	var n1 int64 = n
	r := []int{}
	for n != 0 {
		r = append(r, int(n%10))
		n = n / 10
	}
	n = n1
	var result int64

	for i := 0; i < len(r); i++ {
		var sum int64
		for i := range r {
			sum = sum + int64(r[i])
		}
		if sum <= int64(target) {
			return result
		}
		var cc int64 = 1
		for i := range r {
			if r[i] == 0 {
				cc = cc * 10
			} else {
				cc = cc * int64(10-r[i])
				break
			}
		}
		//        fmt.Println(n, cc)

		n = n + int64(cc)
		n1 = n
		r = []int{}
		for n != 0 {
			r = append(r, int(n%10))
			n = n / 10
		}
		n = n1
		result = result + cc
	}
	return result
}
