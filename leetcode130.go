/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-12 13:21:25
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-12 13:38:29
 */
package myleetcode

func solve(board [][]byte) {
	mark := make([][]int, len(board))
	for y := range mark {
		mark[y] = make([]int, len(board[0]))
	}
	for x := 0; x < len(board[0]); x++ {
		dfs(x, 0, board, mark)
		dfs(x, len(board)-1, board, mark)
	}
	for y := 0; y < len(board); y++ {
		dfs(0, y, board, mark)
		dfs(len(board[0])-1, y, board, mark)
	}
	for y := range board {
		for x := range board[y] {
			if mark[y][x] != 1 && board[y][x] == 'O' {
				board[y][x] = 'X'
			}
		}
	}
}

func dfs(x, y int, board [][]byte, mark [][]int) {
	if y < 0 || y >= len(board) || x < 0 || x >= len(board[0]) || mark[y][x] == 1 || board[y][x] != 'O' {
		return
	}
	mark[y][x] = 1
	dfs(x+1, y, board, mark)
	dfs(x-1, y, board, mark)
	dfs(x, y+1, board, mark)
	dfs(x, y-1, board, mark)
}
