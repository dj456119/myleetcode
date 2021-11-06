/*
 * @Descripttion:编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-06 11:10:45
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-06 11:12:14
 */

package myleetcode

func hammingWeight(num uint32) int {
	count := 0
	for {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
		if num == 0 {
			return count
		}
	}
}

func HammingWeight(num uint32) int {
	return hammingWeight(num)
}
