/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-13 23:28:13
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-13 23:44:33
 */
package myleetcode

func maxCoins(nums []int) int {
	sum := 0
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	dparr := nums
	for {
		if len(nums) == 2 {
			return sum
		}
		if len(nums) == 3 {
			return sum + nums[2]
		}
		if len(nums) == 4 {
			if nums[1] > nums[2] {
				return nums[1]*nums[2] + nums[1]
			}
			return nums[1]*nums[2] + nums[2]
		}
		max := -1
		maxIndex := -1
		for i := range nums {
			if i == 0 {
				continue
			}
			if i == len(nums)-1 {
				break
			}
			z := nums[i-1] * nums[i] * nums[i+1]
			if z > max {
				max = z
				maxIndex = i
			}
		}
		sum = sum + max
		dparr = deleteArr(maxIndex, dparr)
	}
}

func deleteArr(index int, dparr []int) []int {
	for i := range dparr {
		if i == len(dparr)-1 {
			break
		}
		if i >= index {
			dparr[i] = dparr[i+1]
		}
	}
	return dparr[:len(dparr)-1]
}
