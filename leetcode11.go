/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-01 14:20:34
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-01 14:23:47
 */
package myleetcode

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	max := 0
	for {
		h := 0
		if height[i] > height[j] {
			h = height[j]
		} else {
			h = height[i]
		}
		temp := h * (j - i)
		if temp > max {
			max = temp
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
		if i >= j {
			break
		}
	}
	return max
}
