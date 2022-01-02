/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-02 11:43:39
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-02 11:56:34
 */
package myleetcode

func numberOfBeams(bank []string) int {
	prevCountHave := 0
	result := 0
	if len(bank) == 0 || len(bank) == 1 {
		return 0
	}
	for _, j := range bank {
		n := getCount(j)
		if n == 0 {
			continue
		} else {
			result = result + n*prevCountHave
			prevCountHave = n
		}
	}
	return result
}

func getCount(b string) int {
	result := 0
	for _, j := range []byte(b) {
		if j == '1' {
			result++
		}
	}
	return result
}
