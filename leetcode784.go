/*
 * @Descripttion:给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-06 00:34:48
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-06 00:50:14
 */

package myleetcode

import "container/list"

func letterCasePermutation(s string) []string {
	resultList := list.New()
	letterCasePermutationExe(0, len(s), []rune(s), []rune{}, resultList)
	result := []string{}
	for resultList.Len() != 0 {
		result = append(result, string(resultList.Front().Value.([]rune)))
		resultList.Remove(resultList.Front())
	}
	return result
}

func letterCasePermutationExe(n, length int, arr, result []rune, resultArr *list.List) {
	if n == length {
		resultArr.PushBack(result)
		return
	}
	if arr[n] >= '0'-0 && arr[n] <= '9' {
		result = append(result, arr[n])
		letterCasePermutationExe(n+1, length, arr, result, resultArr)
	} else {
		temp := make([]rune, len(result))
		copy(temp, result)
		if arr[n] >= 'a' && arr[n] <= 'z' {
			result = append(result, arr[n])
			temp = append(temp, arr[n]-32)
		} else {
			result = append(result, arr[n])
			temp = append(temp, arr[n]+32)
		}
		letterCasePermutationExe(n+1, length, arr, result, resultArr)
		letterCasePermutationExe(n+1, length, arr, temp, resultArr)
	}

}

func LetterCasePermutation(s string) []string {
	return letterCasePermutation(s)
}
