/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-11-22 00:48:05
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-22 00:48:05
 */
package myleetcode

func decodeCiphertext(encodedText string, rows int) string {
	if rows == 1 {
		return encodedText
	}
	rect := make([][]byte, rows)
	var cs int
	cs = len(encodedText) / rows
	if len(encodedText)%rows != 0 {
		cs++
	}

	for i := range rect {
		rect[i] = make([]byte, cs)
	}
	for i, j := range encodedText {
		y := i / cs
		x := i % cs
		rect[y][x] = byte(j)
	}
	result := []byte{}
	for i := 0; i < cs; i++ {
		x := i
		y := 0
		for x < cs && y < rows {
			result = append(result, rect[y][x])
			x++
			y++
			if x == cs && y < rows {
				return myTrim(result)
			}
		}
	}
	return myTrim(result)

}

func myTrim(result []byte) string {
	count := 0
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] == ' ' {
			count++
		} else {
			break
		}
	}
	return string(result[0 : len(result)-count])
}
