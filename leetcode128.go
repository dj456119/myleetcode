/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-24 23:56:16
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-25 00:47:15
 */
package myleetcode

func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	z := make(map[int]bool)
	for _, j := range nums {
		z[j] = true
	}
	count := 0
	for k := range z {
		if !z[k] {
			continue
		}
		q := k + 1
		m := 0
		_, ok := z[q]
		for ok {
			m++
			z[q] = false
			q++
			_, ok = z[q]
		}
		if m > count {
			count = m
		}
	}
	return count + 1
}
