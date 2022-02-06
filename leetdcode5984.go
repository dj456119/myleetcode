/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-05 23:53:43
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-05 23:57:45
 */
package myleetcode

func minimumSum(num int) int {
	z := [4]int{}
	z1 := z[:]
	z1[0] = num % 10
	z1[1] = num / 10 % 10
	z1[2] = num / 100 % 10
	z1[3] = num / 1000 % 10
	min := 999999
	for i := 0; i < 4; i++ {
		for j := i; j < 4; j++ {
			for m := j; m < 4; m++ {
				for n := m; m < 4; m++ {
					p := z1[m]*10 + z1[n] + z1[i]*10 + z1[j]
					if p < min {
						min = p
					}
				}
			}
		}
	}
	return min
}
