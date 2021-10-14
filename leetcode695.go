/*
 * @Descripttion:给你一个大小为 m x n 的二进制矩阵 grid 。

岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

岛屿的面积是岛上值为 1 的单元格的数目。

计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

 * @version:
 * @Author: cm.d
 * @Date: 2021-10-14 17:23:53
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-14 17:40:15
*/

package myleetcode

func maxAreaOfIsland(grid [][]int) int {
	flag := make([][]int, len(grid))
	for i, j := range grid {
		flag[i] = make([]int, len(j))
	}
	result := 0
	for i := range grid {
		for j := range grid[i] {
			if flag[i][j] == 0 && grid[i][j] == 1 {
				nr := maxAreaOfIslandByFlag(grid, flag, i, j)
				if nr > result {
					result = nr
				}
			}
		}
	}
	return result
}

func maxAreaOfIslandByFlag(grid [][]int, flag [][]int, rs, rc int) int {
	if rs < 0 || rc < 0 || rs >= len(grid) || rc >= len(grid[rs]) {
		return 0
	}
	if flag[rs][rc] == 1 {
		return 0
	}
	if grid[rs][rc] == 0 {
		return 0
	}
	flag[rs][rc] = 1

	result := 0
	result = maxAreaOfIslandByFlag(grid, flag, rs-1, rc) + result
	result = maxAreaOfIslandByFlag(grid, flag, rs+1, rc) + result
	result = maxAreaOfIslandByFlag(grid, flag, rs, rc-1) + result
	result = maxAreaOfIslandByFlag(grid, flag, rs, rc+1) + result

	return result + 1
}

func MaxAreaOfIsland(grid [][]int) int {
	return maxAreaOfIsland(grid)
}
