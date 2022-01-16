/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 10:49:11
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 16:39:48
 */
package myleetcode

func minMoves5194(target int, maxDoubles int) int {
	jiabei := 0
	yidong := 0
	for target != 1 {
		if maxDoubles != 0 {
			if target%2 == 1 {
				target--
				yidong++
				target = target / 2
				jiabei++
				maxDoubles--
			} else {
				target = target / 2
				jiabei++
				maxDoubles--
			}
		} else {
			yidong = yidong + target - 1
			break
		}
	}
	return jiabei + yidong
}
