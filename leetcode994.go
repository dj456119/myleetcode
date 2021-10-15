/*
 * @Descripttion:在给定的网格中，每个单元格可以有以下三个值之一：

值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。

返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。

 * @version:
 * @Author: cm.d
 * @Date: 2021-10-15 17:58:18
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-15 18:59:33
*/
package myleetcode

import "container/list"

func orangesRotting(grid [][]int) int {
	result := 0
	flag := make([][]int, len(grid))
	resultArr := make([][]int, len(grid))

	for i, j := range grid {
		flag[i] = make([]int, len(j))
		resultArr[i] = make([]int, len(j))
	}

	ll := list.New()

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
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
		value := resultArr[x][y]
		rott(grid, resultArr, flag, value, x+1, y, ll)
		rott(grid, resultArr, flag, value, x-1, y, ll)
		rott(grid, resultArr, flag, value, x, y+1, ll)
		rott(grid, resultArr, flag, value, x, y-1, ll)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return -1
			}
			if resultArr[i][j] > result {
				result = resultArr[i][j]
			}
		}
	}

	return result
}

func rott(grid, resultArry, flag [][]int, value, x, y int, ll *list.List) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
		return
	}
	if flag[x][y] == 1 {
		return
	}
	if grid[x][y] == 0 || grid[x][y] == 2 {
		return
	}

	grid[x][y] = 2
	flag[x][y] = 1
	resultArry[x][y] = value + 1
	ll.PushBack([]int{x, y})
}

func OrangesRotting(grid [][]int) int {
	return orangesRotting(grid)
}
