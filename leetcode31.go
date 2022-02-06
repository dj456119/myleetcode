/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-08 17:12:46
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-06 00:44:36
 */
package myleetcode

func nextPermutation(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	a := len(nums) - 1
	b := len(nums) - 2
	min := 99999
	minIndex := -1
	for {
		if b < 0 {
			i := 0
			j := len(nums) - 1
			for i < j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
			return
		} else {
			if nums[b] < nums[a] {
				for i := len(nums) - 1; i >= b+1; i-- {
					if nums[i] > nums[b] && nums[i] < min {
						min = nums[i]
						minIndex = i
					}
				}
				nums[b], nums[minIndex] = nums[minIndex], nums[b]
				i := b + 1
				j := len(nums) - 1
				for i < j {
					nums[i], nums[j] = nums[j], nums[i]
					i++
					j--
				}
				return
			}
		}
		a--
		b--
	}
}
