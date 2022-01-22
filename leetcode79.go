/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-23 00:57:40
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-23 01:29:31
 */
package myleetcode

func exist(board [][]byte, word string) bool {
	for y := range board {
		for x := range board[y] {
			b := make([][]bool, len(board))
			for i := range b {
				b[i] = make([]bool, len(board[0]))
			}

			if dfs(board, y, x, word, 0, b) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, y, x int, word string, wordPos int, al [][]bool) bool {

	if y >= len(board) || y < 0 || x >= len(board[y]) || x < 0 {
		return false
	}
	if board[y][x] == word[wordPos] {

		if al[y][x] {
			return false
		}
		al[y][x] = true
		if wordPos == len(word)-1 {
			return true
		}
		b1 := dfs(board, y-1, x, word, wordPos+1, al)
		if b1 {
			return true
		}
		b2 := dfs(board, y+1, x, word, wordPos+1, al)
		if b2 {
			return true
		}
		b3 := dfs(board, y, x-1, word, wordPos+1, al)
		if b3 {
			return true
		}
		b4 := dfs(board, y, x+1, word, wordPos+1, al)
		if b4 {
			return true
		}
		al[y][x] = false
	}
	return false
}
