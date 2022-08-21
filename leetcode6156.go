package myleetcode

func minimumRecolors(blocks string, k int) int {
	result := 999999999
	for i := 0; i < len(blocks); i++ {
		if i+k > len(blocks) {
			return result
		}
		m := blocks[i : i+k]
		wcount := 0
		for j := range m {
			if m[j] == 'W' {
				wcount++
			}
		}
		if wcount <= result {
			result = wcount
		}
		//        fmt.Println(m, ccc, wcount, bcount)
	}
	return result
}
