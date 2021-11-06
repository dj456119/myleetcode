/*
 * @Descripttion:颠倒给定的 32 位无符号整数的二进制位。
 * @version:
 * @Author: cm.d
 * @Date: 2021-11-06 11:14:10
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-06 11:53:33
 */

package myleetcode

func reverseBits(num uint32) uint32 {
	var result uint32
	count := 0
	for {
		n := num & 1
		result = result | n

		num = num >> 1
		count = count + 1
		if num == 0 {
			break
		}
		result = result << 1
	}

	for i := 0; i < 32-count; i++ {
		result = result << 1
	}
	return result
}

func ReverseBits(num uint32) uint32 {
	return reverseBits(num)
}
