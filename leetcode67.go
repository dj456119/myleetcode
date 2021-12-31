/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 12:50:23
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 13:20:04
 */
package myleetcode

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	aa := reString([]byte(a))
	bb := reString([]byte(b))
	result := []byte{}
	flag := 0
	for i := range aa {
		ta := int(aa[i] - '0')
		tb := ta
		if i >= len(bb) {
			tb = 0
		} else {
			tb = int(bb[i] - '0')
		}
		tc := ta + tb + flag
		flag = tc / 2
		tc = tc % 2
		result = append(result, byte(tc+'0'))
	}
	if flag == 1 {
		result = append(result, '1')
	}
	result = reString(result)
	return string(result)
}

func reString(result []byte) []byte {
	i := 0
	j := len(result) - 1
	for {
		result[i], result[j] = result[j], result[i]
		i++
		j--
		if i >= j {
			break
		}
	}
	return result
}
