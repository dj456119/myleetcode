/*
 * @Descripttion:给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-18 00:18:12
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-18 00:32:10
 */

package myleetcode

func maxProduct(words []string) int {
	wordsMap := make([]uint32, len(words))
	for i, word := range words {
		var temp uint32
		for _, b := range word {
			temp = temp | (1 << (b - 'a'))
		}
		wordsMap[i] = temp
	}
	max := 0
	for i, j := range words {
		for m, n := range words {
			if i == m {
				continue
			}
			if wordsMap[i]&wordsMap[m] == 0 {
				temp := len(j) * len(n)
				if temp > max {
					max = temp
				}
			}
		}
	}
	return max
}

func MaxProduct(words []string) int {
	return maxProduct(words)
}
