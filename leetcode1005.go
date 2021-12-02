/*
 * @Descripttion:给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：

选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
重复这个过程恰好 k 次。可以多次选择同一个下标 i 。

以这种方式修改数组后，返回数组 可能的最大和 。

 * @version:
 * @Author: cm.d
 * @Date: 2021-12-03 00:36:22
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-03 00:46:41
*/
package myleetcode

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	count := 0
	for _, j := range nums {
		if j < 0 {
			count++
		}
	}
	i := 0
	for {
		if count == 0 || k == 0 {
			break
		}
		nums[i] = -nums[i]
		i++
		count--
		k--
	}
	sort.Ints(nums)
	sum := 0
	for _, j := range nums {
		sum = sum + j
	}
	if k%2 == 0 {
		return sum
	} else {
		return sum - 2*nums[0]
	}
}
func LargestSumAfterKNegations(nums []int, k int) int {
	return largestSumAfterKNegations(nums, k)
}
