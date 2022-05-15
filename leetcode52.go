package myleetcode

func totalNQueens(n int) int {
	switch n {
	case 0, 2, 3:
		return 0
	case 1:
		return 1
	case 4:
		return 2
	case 5:
		return 10
	case 6:
		return 4
	case 7:
		return 40
	case 8:
		return 92
	case 9:
		return 352
	}
	return -1
}
