/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 10:57:24
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 11:27:10
 */
package myleetcode

func merge88(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, m+n)
	i, j, z := 0, 0, 0
	for {
		if i == m || j == n {
			break
		}
		if nums1[i] <= nums2[j] {
			result[z] = nums1[i]
			i++
		} else {
			result[z] = nums2[j]
			j++
		}
		z++
	}
	for i < m {

		result[z] = nums1[i]
		z++
		i++
	}
	for j < n {

		result[z] = nums2[j]
		z++
		j++

	}
	copy(nums1, result)
}
