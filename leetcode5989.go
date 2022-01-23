/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-23 10:30:52
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-23 10:41:40
 */
package myleetcode

func countElements(nums []int) int {

	z := make(map[int]int)
	min := nums[0]
	max := nums[0]
	for _, j := range nums {
		if _, ok := z[j]; ok {
			z[j]++
		} else {
			z[j] = 1
		}
		if min > j {
			min = j
		}
		if max < j {
			max = j
		}

	}
	return len(nums) - z[min] - z[max]
}
