/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-20 11:56:53
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-20 12:01:52
 */
package myleetcode

func coutPairs(nums []int, k int) int64 {
	z := 0
	for i := range nums {
		if nums[i]%k == 0 {
			z++
		}
	}
	a1 := len(nums) - z
	sum := int64(z)*int64(a1) + int64(z)*(int64(z)-1)/2
	return sum
}
