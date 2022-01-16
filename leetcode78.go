/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 17:24:58
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 17:29:19
 */
package myleetcode

func subsets(nums []int) [][]int {
	return dfs(nums, 0, []int{})
}

func dfs(nums []int, now int, result []int) [][]int {
	result1 := make([]int, len(result))
	copy(result1, result)
	result1 = append(result1, nums[now])
	now++
	if now == len(nums) {
		return [][]int{result, result1}
	}
	r := [][]int{}
	r1 := dfs(nums, now, result)
	r2 := dfs(nums, now, result1)
	r = append(r, r1...)
	r = append(r, r2...)
	return r
}
