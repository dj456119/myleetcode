/*
 * @Descripttion:给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。

如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-06 10:50:12
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-06 11:10:34
*/

package myleetcode

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}

func IsPowerOfTwo(n int) bool {
	return isPowerOfTwo(n)
}
