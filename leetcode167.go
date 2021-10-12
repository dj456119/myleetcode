/*
 * @Descripttion: leetcode167，有序数组两数之和等于第三个数
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-09 16:33:51
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-12 17:49:19
 */

package myleetcode

func twoSum167(numbers []int, target int) []int {
	if len(numbers) == 0 || len(numbers) == 1 {
		return nil
	}
	start := 0
	end := len(numbers) - 1
	for {
		sum := numbers[start] + numbers[end]
		if sum > target {
			end--
		} else if sum < target {
			start++
		} else {
			return []int{start + 1, end + 1}
		}
		if start >= end {
			return nil
		}
	}
}

func TwoSum167(numbers []int, target int) []int {
	return twoSum167(numbers, target)
}
