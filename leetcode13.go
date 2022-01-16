/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-10 22:30:40
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-10 22:53:14
 */
package myleetcode

func romanToInt(s string) int {
	sum := 0
	zmap := make(map[byte]int)
	zmap['I'] = 1
	zmap['V'] = 5
	zmap['X'] = 10
	zmap['L'] = 50
	zmap['C'] = 100
	zmap['D'] = 500
	zmap['M'] = 1000

	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && zmap[s[i]] < zmap[s[i+1]] {
			sum = sum + zmap[s[i+1]] - zmap[s[i]]
			i++
		} else {
			sum = sum + zmap[s[i]]
		}

	}
	return sum
}
