/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 10:35:22
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 10:38:53
 */

package myleetcode

func divideString(s string, k int, fill byte) []string {
	result := []string{}
	buff := []byte{}
	z := 0
	for _, j := range []byte(s) {
		if z != k {
			buff = append(buff, j)
			z++
		} else {
			z = 0
			result = append(result, string(buff))
			buff = []byte{}
		}
	}
	s1 := []byte(result[len(result)-1])
	for len(s1) != k {
		s1 = append(s1, fill)
	}
	result[len(result)-1] = string(s1)
	return result
}
