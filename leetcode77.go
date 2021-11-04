/*
 * @Descripttion:给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-03 20:54:21
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-04 20:27:03
*/
package myleetcode

import (
	"container/list"
)

func combine(n int, k int) [][]int {
	resultArr := list.New()
	combineArr(1, k, n, 0, []int{}, resultArr)
	resultSlince := make([][]int, resultArr.Len())
	length := resultArr.Len()
	for i := 0; i < length; i++ {
		resultSlince[i] = resultArr.Front().Value.([]int)
		resultArr.Remove(resultArr.Front())
	}
	return resultSlince
}

func combineArr(n, k, j, i int, result []int, resultArr *list.List) {
	if i == k {
		resultArr.PushBack(result)
		return
	}
	if n == j+1 {
		return
	}
	temp := make([]int, len(result))
	copy(temp, result)
	temp1 := append(temp, n)
	combineArr(n+1, k, j, i, result, resultArr)
	combineArr(n+1, k, j, i+1, temp1, resultArr)
}

func Combine(n int, k int) [][]int {
	return combine(n, k)
}
