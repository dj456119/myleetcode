/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-13 10:46:46
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-13 11:37:32
 */
package myleetcode

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	gz := make([][]int, n)
	for y := range gz {
		gz[y] = make([]int, n)
	}
	for m := range artifacts {
		r1 := artifacts[m][0]
		c1 := artifacts[m][1]
		r2 := artifacts[m][2]
		c2 := artifacts[m][3]
		if r1 == r2 {
			for i := c1; i <= c2; i++ {
				gz[i][r1] = m + 1
			}
		} else {
			if c1 == c2 {
				for i := r1; i <= r2; i++ {
					gz[c1][i] = m + 1
				}
			} else {
				for i := r1; i <= r2; i++ {
					for j := c1; j <= c2; j++ {
						gz[j][i] = m + 1
					}

				}

			}
		}
	}

	notin := make(map[int]bool)
	for z := range dig {
		gz[dig[z][1]][dig[z][0]] = 0
	}

	for y := range gz {
		for x := range gz[y] {
			if gz[y][x] != 0 {
				notin[gz[y][x]] = true
			}
		}
	}

	return len(artifacts) - len(notin)
}
