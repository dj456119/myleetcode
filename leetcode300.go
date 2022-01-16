/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-15 19:27:45
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-15 19:46:36
 */
package myleetcode

func lengthOfLIS(nums []int) int {
	arr := make([]int, len(nums))
	result := 0
	for x := range arr {
		if x == 0 {
			arr[0] = 1
			continue
		}

		z := x
		max := 0
		for z >= 0 {
			if nums[x] > nums[z] {
				if arr[z] > max {
					max = arr[z]
				}
			}
			z--
		}
		arr[x] = max + 1
		if arr[x] > result {
			result = arr[x]
		}
	}
	return arr[len(arr)-1]
}
