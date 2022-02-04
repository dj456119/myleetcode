/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-03 19:07:04
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-04 22:39:33
 */
package myleetcode

func subarraySum(nums []int, k int) int {
	dparr := make(map[int]int)
	sum := 0
	for _, j := range nums {
		dparr[sum]++
		sum = sum + j
	}
	result := 0
	for dpk, dpv := range dparr {
		if dpk == k {
			result++
		}
		if dpvv, ok := dparr[dpk-k]; ok {
			result = result + dpvv*dpv
			dparr[dpk-k] = 0
			dparr[dpk] = 0
		}
	}
	return result
}
