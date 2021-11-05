/*
 * @Descripttion:给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
 * @version:
 * @Author: cm.d
 * @Date: 2021-11-04 23:37:04
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-04 23:54:11
 */
package myleetcode

import "container/list"

func permute(nums []int) [][]int {
	resultArr := list.New()
	permuteExe(0, len(nums), nums, []int{}, resultArr)
	resultSlince := make([][]int, resultArr.Len())
	length := resultArr.Len()
	for i := 0; i < length; i++ {
		resultSlince[i] = resultArr.Front().Value.([]int)
		resultArr.Remove(resultArr.Front())
	}
	return resultSlince
}

func permuteExe(pos, length int, nums, result []int, resultArr *list.List) {
	for i, j := range nums {
		temp := []int{}
		for m, n := range nums {
			if m != i {
				temp = append(temp, n)
			}
		}
		resultTemp := make([]int, len(result))
		copy(resultTemp, result)
		resultTemp = append(resultTemp, j)
		if pos+1 == length {
			resultArr.PushBack(resultTemp)
			return
		}
		permuteExe(pos+1, length, temp, resultTemp, resultArr)
	}
}

func Permute(nums []int) [][]int {
	return permute(nums)
}
