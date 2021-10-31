/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-31 10:03:35
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-31 13:58:42
 */
package myleetcode

type voidStruct struct {
}

type void voidStruct

func findWords(words []string) []string {
	result := []string{}
	var a void
	mapOne := make(map[byte]void)
	mapTwo := make(map[byte]void)
	mapThree := make(map[byte]void)
	for _, j := range "qwertyuiop" {
		mapOne[byte(j)] = a
	}
	for _, j := range "asdfghjkl" {
		mapTwo[byte(j)] = a
	}
	for _, j := range "zxcvbnm" {
		mapThree[byte(j)] = a
	}
	for _, word := range words {
		isOk := true
		firstWord := byte(word[0])
		flag := getNum(firstWord, mapOne, mapTwo, mapThree)
		for _, b := range word {
			if flag != getNum(byte(b), mapOne, mapTwo, mapThree) {
				isOk = false
				break
			}
		}
		if isOk {
			result = append(result, word)
		}

	}
	return result
}

func getNum(b byte, mapOne, mapTwo, mapThree map[byte]void) int {
	if b < 97 {
		b = b + 32
	}
	if _, ok := mapOne[byte(b)]; ok {
		return 1
	}
	if _, ok := mapTwo[byte(b)]; ok {
		return 2
	}
	if _, ok := mapThree[byte(b)]; ok {
		return 3
	}
	return -1
}

func FindWords(words []string) []string {
	return findWords(words)
}
