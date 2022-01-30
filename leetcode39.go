/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 20:00:30
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-30 00:24:51
 */
package myleetcode

func combinationSum(candidates []int, target int) [][]int {
	return dfs39(candidates, 0, 0, target, []int{})
}

func dfs39(candidates []int, now int, sum int, target int, result []int) [][]int {
	if now == len(candidates) {
		return nil
	}

	rr := [][]int{}
	r1 := []int{}
	for sum <= target {
		r2 := make([]int, len(result))
		copy(r2, result)
		r2 = append(r2, r1...)
		if sum == target {
			rr = append(rr, r2)
			break
		}
		rrr := dfs39(candidates, now+1, sum, target, r2)
		if rrr != nil {
			rr = append(rr, rrr...)
		}
		sum = sum + candidates[now]
		r1 = append(r1, candidates[now])
	}

	return rr
}
