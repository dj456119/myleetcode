/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-14 21:55:57
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-06 21:54:51
 */
package myleetcode

func kthSmallest(matrix [][]int, k int) int {
	start := matrix[0][0]
	end := matrix[len(matrix)-1][len(matrix)-1]
	for {
		z := get(matrix, (start+end)/2)
		if z == k-1 {
			return z
		} else if z > k-1 {
			end = z
		} else {
			start = z + 1
		}
		if start >= end {
			break
		}
	}
	return -1
}

func get(matrix [][]int, k int) int {
	i := 0
	j := 0
	count := 0
	for {
		if matrix[i][j] > k {
			i--
			break
		}
		i++
	}
	count = count + i + 1
	for {
		j++
		if j >= len(matrix[i]) || i < 0 {
			break
		}
		for {
			if matrix[i][j] > k {
				i--
			} else {
				break
			}
		}
		count = count + i + 1
	}
	return count
}
