/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-30 10:53:34
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-30 14:04:29
 */
package myleetcode

import (
	"math"
)

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	arr := make([][]int, len(s))
	for i := range arr {
		arr[i] = make([]int, len(arr))
	}

	temp0 := make([]int, len(s))
	for i, j := range []byte(s) {
		temp0[i] = int(j-'a') + 1
		if k == 1 {
			if temp0[i]%modulo == hashValue {
				return string(j)
			}
		}
	}

	for y := range arr {
		if y == 0 {
			p := 1
			sum := 0
			if y+k >= len(s)+1 {
				continue
			}
			for x := 0; x < k; x++ {
				hashValuez := temp0[x] * p
				p = p * power
				sum = sum + hashValuez
				arr[y][x] = sum
			}
			if sum%modulo == hashValue {
				return s[y : y+k]
			}
		} else {
			if y+k >= len(s)+1 {
				continue
			}
			z := arr[y-1][y+k-2] - temp0[y-1]
			// fmt.Println(y, z,temp0[y+k-1],int(math.Pow(float64(power), float64(k-1))) )
			z = z/power + temp0[y+k-1]*int(math.Pow(float64(power), float64(k-1)))
			arr[y][y+k-1] = z
			if z%modulo == hashValue {
				return s[y : y+k]
			}
		}
		//   fmt.Println(arr[y])
	}
	return ""
}
