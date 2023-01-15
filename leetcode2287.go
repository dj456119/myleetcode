package myleetcode

func rearrangeCharacters2287(s string, target string) int {
	z := make(map[byte]int)
	z1 := make(map[byte]int)
	for i := range target {
		z[target[i]]++
	}
	for j := range s {
		z1[s[j]]++
	}
	result := 66666666666
	for k, v := range z {
		if z1[k] == 0 {
			return 0
		}
		q := z1[k] / v
		if q == 0 {
			return 0
		}
		if result > q {
			result = q
		}
	}
	return result
}
