/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-01 14:26:17
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-01 14:36:15
 */
package myleetcode

func canJump(nums []int) bool {
	nowPos := len(nums) - 1
	for {
		for i := nowPos - 1; i >= 0; i-- {
			if (nowPos - i) <= nums[i] {
				nowPos = i
				break
			}
			if i == 0 && nowPos != 0 {
				return false
			}
		}
		if nowPos == 0 {
			return true
		}
	}
}
