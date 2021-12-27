/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-26 10:56:29
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-26 11:04:42
 */
package main

import "math"

func getDistances(arr []int) []int64 {
	tempMap := make(map[int][]int)
	result := make([]int64, len(arr))
	for i, j := range arr {
		if arrtemp, ok := tempMap[j]; ok {
			arrtemp = append(arrtemp, i)
			tempMap[j] = arrtemp
		} else {
			arrtemp := []int{}
			arrtemp = append(arrtemp, i)
			tempMap[j] = arrtemp
		}
	}
	for i, j := range arr {
		arrtemp := tempMap[j]
		var sum int64
		for _, m := range arrtemp {
			sum = sum + int64(math.Abs(float64(i-m)))
		}
		result[i] = sum
	}
	return result
}
