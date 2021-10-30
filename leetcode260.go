/*
 * @Descripttion:给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。


 * @version:
 * @Author: cm.d
 * @Date: 2021-10-30 10:17:49
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-30 10:23:05
 */

package myleetcode

func singleNumber(nums []int) []int {
	result := []int{}
	resultMap := make(map[int]int)
	for _, j := range nums {
		if count, ok := resultMap[j]; ok {
			resultMap[j] = count + 1
		} else {
			resultMap[j] = 1
		}
	}

	for k, v := range resultMap {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}

func SingleNumber(nums []int) []int {
	return singleNumber(nums)
}
