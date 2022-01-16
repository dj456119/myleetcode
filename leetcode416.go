/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-15 13:01:08
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-15 13:24:59
 */
package myleetcode

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum := 0
	max := 0
	for _, j := range nums {
		sum = sum + j
		if j > max {
			max = j
		}
	}
	if sum%2 != 0 {
		return false
	}
	if max > sum/2 {
		return false
	}
	target := sum / 2
	arr := make([][]bool, len(nums))
	for i := range arr {
		arr[i] = make([]bool, target+1)
	}

	for y, z := range arr {
		for x := range z {
			if y == 0 {
				if x == nums[y] {
					z[x] = true
				}
			} else {
				if x == nums[y] {
					z[x] = true
				} else {
					if arr[y-1][x] == true {
						z[x] = true
					} else {
						if arr[y-1][x-nums[y]] == true {
							z[x] = true
						}
					}
				}
			}
			if arr[y][target] {
				return true
			}
		}
	}
	return false
}
