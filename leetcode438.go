/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-03 18:34:29
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-03 18:44:49
 */
package myleetcode

func findAnagrams(s string, p string) []int {
	pArr := make([]int, 26)
	sArr := make([]int, 26)
	if len(s) < len(p) {
		return []int{}
	}
	for _, j := range []byte(p) {
		pArr[j-'a']++
	}
	i := 0
	j := len(p) - 1
	for m := i; m <= j; m++ {
		sArr[s[m]-'a']++
	}
	result := []int{}

	for {
		if compare(pArr, sArr) {
			result = append(result, i)
		}
		i++
		j++
		if j >= len(s) {
			break
		}
		sArr[s[j]-'a']++
		sArr[s[i-1]-'a']--
	}
	return result
}

func compare(pArr, sArr []int) bool {
	for i := range pArr {
		if pArr[i] == sArr[i] {
			if i == len(pArr)-1 {
				return true
			}
		} else {
			break
		}
	}
	return false
}
