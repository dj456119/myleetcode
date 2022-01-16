/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 20:00:30
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 20:41:48
 */
package myleetcode

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	return digui(candidates, 0, target, []int{})
}

func digui(candidates []int, now int, target int, result []int) [][]int {
	fmt.Println(result)
	if now == len(candidates) {
		if target == 0 {
			return [][]int{result}
		}
		return [][]int{}
	}

	if now == len(candidates)-1 {
		if candidates[now] > target {
			return [][]int{}
		}
	}

	fin := [][]int{}
	z := digui(candidates, now+1, target, result)
	if len(z) != 0 {
		fin = append(fin, z...)
	}
	r := make([]int, len(result))
	copy(r, result)
	for {
		if target-candidates[now] >= 0 {
			r = append(r, candidates[now])
			r1 := make([]int, len(r))
			copy(r1, r)
			fmt.Println("shenmi:", r, target, now)
			z := digui(candidates, now+1, target-candidates[now], r1)
			if len(z) != 0 {
				fin = append(fin, z...)
			}
			target = target - candidates[now]
		} else {
			break
		}
	}
	if target > 0 {
		r := make([]int, len(result))
		z := digui(candidates, now+1, target, r)
		if len(z) != 0 {
			fin = append(fin, z...)
		}
		//fmt.Println(fin)
		//break
	}
	return fin
}
