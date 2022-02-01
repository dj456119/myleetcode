/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-01 09:56:02
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-01 10:29:22
 */
package myleetcode

func numIslands(grid [][]byte) int {
	result := 0
	temp := make([][]bool, len(grid))
	for i := range temp {
		temp[i] = make([]bool, len(grid[i]))
	}
	// true已搜索 false未搜索
	for i := range grid {
		for j := range grid[i] {
			if temp[i][j] == false {
				if grid[i][j] == '0' {
					temp[i][j] = true
					continue
				}
				if bfs(grid, temp, j, i) {
					//   fmt.Println(j,i)
					result++
				}
			}
		}
	}
	return result
}

func bfs(grid [][]byte, temp [][]bool, x, y int) bool {
	mylist := [][]int{}
	mylist = append(mylist, []int{x, y})
	for len(mylist) != 0 {
		//   fmt.Println(mylist)
		ele := mylist[0]
		mylist = mylist[1:]
		y = ele[1]
		x = ele[0]
		if temp[y][x] == false {
			temp[y][x] = true
			if grid[y][x] == '0' {
				continue
			}
			if x >= 1 {
				mylist = append(mylist, []int{x - 1, y})
			}
			if x < len(grid[y])-1 {
				mylist = append(mylist, []int{x + 1, y})
			}
			if y < len(grid)-1 {
				mylist = append(mylist, []int{x, y + 1})
			}
			if y >= 1 {
				mylist = append(mylist, []int{x, y - 1})
			}
		}
	}
	return true
}
