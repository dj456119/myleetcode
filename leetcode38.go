/*
 * @Descripttion:给定一个正整数 n ，输出外观数列的第 n 项。

「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。

你可以将其视作是由递归公式定义的数字字符串序列：

countAndSay(1) = "1"
countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。

 * @version:
 * @Author: cm.d
 * @Date: 2021-10-15 19:03:04
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-15 20:12:48
*/
package myleetcode

import "fmt"

func countAndSay(n int) string {
	var result string = "1"
	for i := 1; i < n; i++ {
		result = countAndSaySingle(result)
	}
	return result
}

func countAndSaySingle(n string) string {
	result := ""
	count := 0
	var last rune = -1
	for _, j := range n {
		if j != last && last != -1 {

			result = fmt.Sprintf("%s%d%c", result, count, last)
			last = j
			count = 1
			continue
		}
		count++
		last = j
	}
	return fmt.Sprintf("%s%d%c", result, count, last)
}

func CountAndSay(n int) string {
	return countAndSay(n)
}
