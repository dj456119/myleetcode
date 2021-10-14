/*
 * @Descripttion:给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。

 * @version:
 * @Author: cm.d
 * @Date: 2021-10-14 19:23:21
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-14 22:02:27
*/

package myleetcode

import "container/list"

func updateMatrix(mat [][]int) [][]int {
	result := make([][]int, len(mat))
	flag := make([][]int, len(mat))

	for i, j := range mat {
		result[i] = make([]int, len(j))
		flag[i] = make([]int, len(j))
	}

	ll := list.New()
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				ll.PushBack([]int{i, j})
				flag[i][j] = 1
			}
		}
	}
	for ll.Len() != 0 {
		pos := ll.Front()
		ll.Remove(pos)
		posArr := pos.Value.([]int)

		x := posArr[0]
		y := posArr[1]

		value := result[x][y]

		BFS(x-1, y, value, result, flag, mat, ll)
		BFS(x+1, y, value, result, flag, mat, ll)
		BFS(x, y-1, value, result, flag, mat, ll)
		BFS(x, y+1, value, result, flag, mat, ll)

	}

	return result
}

func BFS(x, y, value int, result, flag, mat [][]int, ll *list.List) {
	if x < 0 || y < 0 || x >= len(mat) || y >= len(mat[x]) {
		return
	}
	if flag[x][y] == 1 || mat[x][y] == 0 {
		return
	}
	ll.PushBack([]int{x, y})
	if mat[x][y] == 1 {
		result[x][y] = value + 1
	}
	flag[x][y] = 1

}

func UpdateMatrix(mat [][]int) [][]int {
	return updateMatrix(mat)
}
