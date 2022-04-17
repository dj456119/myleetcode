package myleetcode

import "fmt"

func digitSum(s string, k int) string {
	m := []byte(s)
	for k < len(m) {
		m = split(m, k)
	}
	return string(m)
}

func split(s []byte, k int) []byte {
	z := [][]byte{}
	for len(s) >= k {
		z = append(z, s[:k+1])
	}
	if len(s) != 0 {
		z = append(z, s)
	}
	result := []byte{}
	for i := range z {
		sum := 0
		for j := range z[i] {
			sum = sum + int(z[i][j]-'0')
		}
		c := fmt.Sprintf("%d", sum)
		result = append(result, []byte(c)...)
	}
	return result
}
