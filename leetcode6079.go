package myleetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	result := strings.Fields(sentence)
	r2 := []byte{}
	for i := range result {
		r2 = append(r2, ' ')
		b, f := isPrice(result[i])
		if b {
			if discount == 100 {
				f = 0
			} else {
				f = f * float64(float64(100-discount)/100.0)
			}
			ss := fmt.Sprintf("%s%.2f", "$", f)
			r2 = append(r2, []byte(ss)...)
		} else {
			r2 = append(r2, []byte(result[i])...)
		}
	}
	return string(r2[1:])
}

func isPrice(s string) (bool, float64) {
	if s[0] != '$' {
		return false, 0
	}
	f, err := strconv.ParseFloat(s[1:], 64)
	if err != nil {
		return false, 0
	}
	return true, f
}
