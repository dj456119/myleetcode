/*
 * @Descripttion:给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。

请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。

 * @version:
 * @Author: cm.d
 * @Date: 2021-10-26 00:21:55
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-26 01:00:42
*/

package myleetcode

import "container/list"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	num2Map := make(map[int]int)
	stack := list.New()
	for i := len(nums2) - 1; i >= 0; i-- {

		for {
			if stack.Len() == 0 {
				num2Map[nums2[i]] = -1
				break
			}
			topTemp := stack.Back()
			top := topTemp.Value.(int)
			if top <= nums2[i] {
				stack.Remove(topTemp)
			} else {
				num2Map[nums2[i]] = top
				break
			}
		}
		stack.PushBack(nums2[i])
	}
	for i, j := range nums1 {
		result[i] = num2Map[j]
	}
	return result
}

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	return nextGreaterElement(nums1, nums2)
}
