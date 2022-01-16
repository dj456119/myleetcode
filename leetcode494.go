/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-15 15:37:53
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-15 17:10:42
 */
package myleetcode

func findTargetSumWays(nums []int, target int) int {
	arr := make([][]int, len(nums))
	sum := 0
	for _, j := range nums {
		sum = sum + j
	}
	for i := range arr {
		arr[i] = make([]int, 2*sum+1)
	}

	for y := range arr {
		for x := range arr[y] {
			if y == 0 {
				if -nums[0] == x-sum {
					arr[0][x]++
				}
				if nums[0] == x-sum {
					arr[0][x]++
				}

			} else {
				if arr[y-1][x] != 0 {

					arr[y][getPos(sum, x-sum+nums[y])] += arr[y-1][x]

					arr[y][getPos(sum, x-sum-nums[y])] += arr[y-1][x]

				}
			}
		}
	}

	for y := range arr {
		for x := range arr[y] {
			if x-sum == target && y == len(arr)-1 {
				return arr[y][x]
			}
		}
	}
	return 0
}

func getPos(sum int, a int) int {
	z := -sum - a
	if z > 0 {
		return z
	} else {
		return -z
	}
}
