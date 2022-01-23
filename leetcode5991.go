/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-23 10:45:49
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-23 10:47:02
 */
package myleetcode

func rearrangeArray(nums []int) []int {
	r1 := []int{}
	r2 := []int{}
	for _, j := range nums {
		if j > 0 {
			r1 = append(r1, j)
		} else {
			r2 = append(r2, j)
		}
	}
	result := []int{}
	i := 0
	for i < len(r1) {
		result = append(result, r1[i])
		result = append(result, r2[i])
		i++
	}
	return result
}
