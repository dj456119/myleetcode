/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-27 00:26:27
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-27 00:37:23
 */
package myleetcode

func countValidWords(sentence string) int {
	result := 0
	buff := []byte{}
	for _, j := range []byte(sentence) {
		if j == ' ' && len(buff) != 0 {
			b := pd(buff)
			if b {
				result++
			}
			buff = []byte{}
		} else if j != ' ' {
			buff = append(buff, j)
		}
	}
	b := pd(buff)
	if b {
		result++
	}
	return result
}

func pd(buff []byte) bool {
	count1 := 0
	for i, j := range buff {
		if (j >= 'a' && j <= 'z') || (j >= '0') && (j <= 9) || j == '-' || j == '!' || j == '.' || j == ',' {
			if j == '-' {
				if i == 0 || i == len(buff)-1 {
					return false
				}
				z := buff[i-1]
				z1 := buff[i+1]
				if (z >= 'a' && z <= 'z') && (z1 >= 'a' && z1 <= 'z') {
					count1++
					if count1 > 1 {
						return false
					}
				} else {
					return false
				}
			} else if j == '!' || j == '.' || j == ',' {
				if i != len(buff)-1 {
					return false
				}
			}
		} else {
			return false
		}
	}
	return false
}
