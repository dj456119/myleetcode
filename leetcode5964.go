/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-26 11:07:30
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 10:37:13
 */
package myleetcode

func executeInstructions(n int, startPos []int, s string) []int {
	result := []int{}
	for i := 0; i < len(s); i++ {
		r := step(n, startPos, s[i:])
		result = append(result, r)
	}
	return result
}

//"RRDDLU", [1,5,4,3,1,0]
func step(n int, startPos []int, s string) int {
	count := 0
	Y := startPos[0]
	X := startPos[1]
	for _, j := range s {
		count++
		switch j {
		case 'R':
			X = X + 1
		case 'D':
			Y = Y + 1
		case 'U':
			Y = Y - 1
		case 'L':
			X = X - 1
		}
		if X < 0 || X >= n || Y < 0 || Y >= n {
			return count - 1
		}
	}
	return count
}
