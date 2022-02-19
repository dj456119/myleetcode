/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-19 13:12:16
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-19 14:50:40
 */
package myleetcode

func largestRectangleArea(heights []int) int {
	dparr := make([][]int, len(heights))
	for i := range dparr {
		dparr[i] = make([]int, 2)
	}
	stc := []int{}
	for i := range dparr {
		dparr[i][0] = -1
		dparr[i][1] = len(dparr)
		if len(stc) == 0 {
			dparr[i][0] = -1
			stc = append(stc, i)
			continue
		}
		if heights[i] >= heights[stc[len(stc)-1]] {
			dparr[i][0] = stc[len(stc)-1]

		} else {
			for heights[i] < heights[stc[len(stc)-1]] {
				j := stc[len(stc)-1]
				stc = stc[:len(stc)-1]
				dparr[j][1] = i

				if len(stc) == 0 {
					dparr[i][0] = -1
					break
				}
				dparr[i][0] = stc[len(stc)-1]
			}

		}
		stc = append(stc, i)
	}
	max := 0
	for i := range dparr {
		z := (dparr[i][1] - dparr[i][0] - 1) * heights[i]
		if z > max {
			max = z
		}
	}
	return max
}
