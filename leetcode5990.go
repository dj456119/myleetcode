/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-23 10:50:01
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-23 10:54:36
 */
package myleetcode

func findLonely(nums []int) []int {
	z := make(map[int]int)
	result := []int{}
	for _, j := range nums {
		if _, ok := z[j]; ok {
			z[j]++
		} else {
			z[j] = 1
		}
	}
	for _, j := range nums {
		if m, ok := z[j]; ok {
			if m == 1 {
				_, ok1 := z[j+1]
				_, ok2 := z[j-1]
				if ok1 || ok2 == false {
					result = append(result, j)
				}
			}
		}
	}
	return result
}
