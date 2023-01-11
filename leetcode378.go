/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-14 21:55:57
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-06 21:54:51
 */
package myleetcode

import (
	"container/heap"
	"sort"
)

func kthSmallest(matrix [][]int, k int) int {
	z := make([]int, 0)
	h := new(HP378)
	h.H = z
	heap.Init(h)
	for i := range matrix {
		for j := range matrix {
			if len(h.H) < k {
				heap.Push(h, matrix[i][j])
				continue
			}
			if matrix[i][j] < h.H[0] {
				heap.Pop(h)
				heap.Push(h, matrix[i][j])
			}
		}
	}
	return heap.Pop(h).(int)
}

type HP378 struct {
	H sort.IntSlice
}

func (hp HP378) Less(i, j int) bool {
	return hp.H[i] > hp.H[j]
}

func (hp HP378) Len() int {
	return len(hp.H)
}

func (hp HP378) Swap(i, j int) {
	hp.H[i], hp.H[j] = hp.H[j], hp.H[i]
}

func (hp *HP378) Push(x interface{}) {
	hp.H = append(hp.H, x.(int))
}

func (hp *HP378) Pop() interface{} {
	z := hp.H[len(hp.H)-1]
	hp.H = hp.H[:len(hp.H)-1]
	return z
}
