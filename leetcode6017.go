/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-06 11:51:53
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-06 12:30:00
 */
package myleetcode

func minimalKSum(nums []int, k int) int64 {
	k1 := int64(k)
	var sumk int64 = k1 * (k1 + 1) / 2
	zm := make(map[int64]int)
	for i := range nums {
		zm[int64(nums[i])] = 1
	}
	for i := range nums {
		if int64(nums[i]) <= int64(k) && zm[int64(nums[i])] != 0 {
			sumk = sumk - int64(nums[i])
			for {
				k1++
				if _, ok := zm[k1]; ok {
					continue
				}
				sumk = sumk + k1
				zm[int64(nums[i])] = 0
				break
			}
		}
	}
	return sumk
}
