/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-08 23:02:36
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-08 23:50:47
 */
package myleetcode

func capitalizeTitle(title string) string {
	a := []byte{}
	for _, j := range []byte(title) {
		if j >= 'A' && j <= 'Z' {
			a = append(a, j+('a'-'A'))
		} else {
			a = append(a, j)
		}
	}
	length := 0
	index := 0
	for i := range a {
		if a[i] == ' ' {
			if length >= 3 {
				a[index] = a[index] - ('a' - 'A')
				length = 0
				index = i + 1
			} else {
				length = 0
				index = i + 1
			}
		} else {
			length++
		}
	}
	if length >= 3 {
		a[index] = a[index] - ('a' - 'A')
	}
	return string(a)
}
