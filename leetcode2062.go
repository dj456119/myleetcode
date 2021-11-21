/*
 * @Descripttion:子字符串 是字符串中的一个连续（非空）的字符序列。

元音子字符串 是 仅 由元音（'a'、'e'、'i'、'o' 和 'u'）组成的一个子字符串，且必须包含 全部五种 元音。

给你一个字符串 word ，统计并返回 word 中 元音子字符串的数目 。

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-22 00:49:05
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-22 00:49:05
*/
package myleetcode

func countVowelSubstrings(word string) int {
	count := 0
	yyMap := make(map[byte]bool)
	yyMap['a'] = true
	yyMap['e'] = true
	yyMap['i'] = true
	yyMap['o'] = true
	yyMap['u'] = true

	start := 0
	end := 0
	for start = 0; start < len(word)-4; start++ {
		for end = 4; end < len(word); end++ {
			if _, ok := yyMap[word[start]]; !ok {
				continue
			}
			if _, ok := yyMap[word[end]]; !ok {
				continue
			}
			if isVowe(start, end, word, yyMap) {
				count++
			}
		}
	}
	return count
}

func isVowe(start, end int, word string, tempMap map[byte]bool) bool {
	tMap := make(map[byte]bool)
	for i := start; i <= end; i++ {
		if _, ok := tempMap[word[i]]; !ok {
			return false
		} else {
			tMap[word[i]] = true
		}
	}
	return len(tMap) == 5
}
