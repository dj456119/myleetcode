/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-10 00:03:47
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-10 00:14:58
 */
package myleetcode

func shortestCompletingWord(licensePlate string, words []string) string {
	arr := make([]int, 26)
	for _, j := range licensePlate {
		if j >= 'a' && j <= 'z' {
			arr[j-'a']++
		} else {
			if j >= 'A' && j <= 'Z' {
				arr[j-'A']++
			}
		}
	}

	resultLength := -1
	resultString := ""

	for _, word := range words {
		arrTemp := make([]int, 26)
		flag := false
		for _, j := range word {
			if j >= 'a' && j <= 'z' {
				arrTemp[j-'a']++
			} else {
				if j >= 'A' && j <= 'Z' {
					arrTemp[j-'A']++
				}
			}
		}

		for i, j := range arr {
			if j > 0 {
				if arrTemp[i] < j {
					flag = true
					break
				}
			}
		}
		if flag {
			continue
		}
		if resultLength == -1 {
			resultLength = len(word)
			resultString = word
		} else {
			if resultLength > len(word) {
				resultLength = len(word)
				resultString = word
			}
		}
	}

	return resultString
}

func ShortestCompletingWord(licensePlate string, words []string) string {
	return shortestCompletingWord(licensePlate, words)
}
