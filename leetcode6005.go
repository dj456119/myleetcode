/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-13 10:35:10
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-13 10:39:14
 */
package myleetcode

func minimumOperations(nums []int) int {
	z1 := make(map[int]int)
	z2 := make(map[int]int)
	max1 := -9999999
	m1 := 0
	max2 := -9999999
	m2 := 0
	for i := range nums {
		if i%2 == 0 {
			z1[nums[i]]++
			if z1[nums[i]] > max1 {
				m1 = nums[i]
			}
		} else {
			z2[nums[i]]++
			if z2[nums[i]] > max2 {
				m2 = nums[i]
			}
		}
	}
	result := 0
	for i := range nums {
		if i%2 == 0 {
			if nums[i] != m1 {
				result++
			}
		} else {
			if nums[i] != m2 {
				result++
			}
		}
	}
	return result
}
