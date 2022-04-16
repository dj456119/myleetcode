package myleetcode

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var result int64
	for i := 0; ; i++ {
		temp := total - cost1*i
		if temp == 0 {
			return result + 1
		}
		if temp < 0 {
			return result
		}
		c := temp / cost2
		result = result + int64(c)
	}
}
