/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 21:28:59
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 21:33:51
 */
package myleetcode

import "fmt"

func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string][]string)
	for _, str := range strs {
		key := do(str)
		if r, ok := resultMap[key]; ok {
			r = append(r, str)
			resultMap[key] = r
		} else {
			resultMap[key] = []string{str}
		}
	}
	result := [][]string{}
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}

func do(s string) string {
	arr := make([]int, 26)
	for _, j := range []byte(s) {
		arr[j-'a']++
	}
	key := "'"
	for _, z := range arr {
		key = fmt.Sprintf("%s_%d", key, z)
	}
	return key
}
