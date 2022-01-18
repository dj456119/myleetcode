/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-19 00:42:44
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-19 01:34:22
 */
package myleetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	z := make(map[int]bool)
	first := nums[0]
	i := 0
	j := 0
	flag := false
	for {

		if j == len(nums) {
			return false
		}
		if len(z) >= k+1 {
			delete(z, first)
			flag = true
		}

		if _, ok := z[nums[j]]; ok {
			return true
		} else {
			z[nums[j]] = true
		}
		if flag {
			i++
			first = nums[i]
		}
		j++
	}
}
