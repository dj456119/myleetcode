/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 17:35:10
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 17:47:19
 */
package myleetcode

func rob(nums []int) int {
	arrno1 := make([]int, len(nums))
	arrin1 := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			arrno1[0] = 0
			arrin1[0] = nums[0]

		} else if i != len(nums)-1 {
			arrno1[i] = max(arrno1[i-1], arrno1[i-2]+nums[i])
			arrin1[i] = max(arrin1[i-1], arrin1[i-2])
		} else if i == 1 {
			arrno1[1] = nums[1]
			arrin1[1] = max(nums[0], nums[1])
		} else {
			arrno1[i] = max(arrno1[i-1], arrno1[i-2]+nums[i])
			arrin1[i] = max(arrin1[i-1], arrin1[i-2]+nums[i])
		}
	}
	return max(arrin1[len(arrin1)-1], arrno1[len(arrno1)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
