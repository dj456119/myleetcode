/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 11:21:46
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 11:49:57
 */
package myleetcode

func mostPoints(questions [][]int) int64 {
	arr := make([][]int64, len(questions))
	temp := make(map[int]int64)
	for i := range arr {
		arr[i] = make([]int64, len(questions))
	}
	var result int64 = 0
	for y := range arr {
		jump := questions[y][1]
		var z int64 = int64(questions[y][0])
		if max, ok := temp[y]; ok {
			z = max
		} else {
			temp[y] = z
			if z > result {
				result = z
			}
		}
		for x := y + jump + 1; x < len(arr); x++ {
			z1 := int64(questions[x][0]) + z
			if z1 > z {
				temp[x] = z
				if z > result {
					result = z
				}
			}
		}
	}
	return result
}

func max(a1, a2 int) int {
	if a1 > a2 {
		return a1
	}
	return a2
}
