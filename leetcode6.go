/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-06 23:47:06
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-07 00:26:20
 */
package myleetcode

func convert(s string, numRows int) string {

	pad := make([][]byte, numRows)
	z := numRows + (numRows) - 1 - 1
	if z == 0 {
		return s
	}
	if len(s) <= numRows {
		return s
	}

	z1 := len(s) / z
	if z1 == 0 {
		z1 = 1
	}
	l := 0

	l = (z1 + 1) * (numRows)

	for i := range pad {
		pad[i] = make([]byte, l)
	}
	x, y := 0, 0
	i := 0
	for i != len(s) {
		for y != numRows-1 {
			if i >= len(s) {
				break
			}
			pad[y][x] = s[i]
			i++
			y++
		}
		for y != 0 {
			if i >= len(s) {
				break
			}
			pad[y][x] = s[i]
			i++
			y--
			x++
		}
	}
	resultbyte := []byte{}
	for _, j := range pad {
		for _, m := range j {
			if m != 0 {
				resultbyte = append(resultbyte, m)
			}
		}
	}
	return string(resultbyte)
}
