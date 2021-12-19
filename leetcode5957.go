/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-19 12:45:12
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-19 12:45:12
 */
package myleetcode

func addSpaces(s string, spaces []int) string {
	result := make([]byte, len(s)+len(spaces))
	mapTemp := make(map[int]bool)
	for _, j := range spaces {
		mapTemp[j] = true
	}
	m := 0
	for i, j := range []byte(s) {
		if _, ok := mapTemp[i]; ok {
			result[m] = ' '
			m++
		}
		result[m] = j
		m++
	}
	return string(result)
}
