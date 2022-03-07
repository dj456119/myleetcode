/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-05 22:44:48
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-05 22:51:54
 */
package myleetcode

func mostFrequent(nums []int, key int) int {
	max := -1
	maxi := 0
	for i := range nums {
		if nums[i] == key {
			if i == len(nums)-1 {
				break
			}
			target := nums[i+1]
			count := 0
			for z := i + 1; z < len(nums); z++ {
				if nums[z] == target {
					count++
				}

			}
			if count > max {
				max = count
				maxi = target
			}
		}
	}
	return maxi
}
