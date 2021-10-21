/*
 * @Descripttion:给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

 * @version:
 * @Author: cm.d
 * @Date: 2021-10-21 14:19:38
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-21 14:54:48
*/

package myleetcode

func plusOne(digits []int) []int {

	flag := false

	for i := range digits {
		if digits[i] != 9 {
			flag = true
		}
	}

	if !flag {
		result := make([]int, len(digits)+1)
		for i := range digits {
			result[i+1] = 0
		}
		result[0] = 1
		return result
	} else {
		aFlag := true
		result := make([]int, len(digits))
		for i := len(digits) - 1; i >= 0; i-- {
			if !aFlag {
				digits[i] = digits[i] + 1
				aFlag = true
			}
			if i == len(digits)-1 {
				digits[i] = digits[i] + 1
			}

			if digits[i] >= 10 {
				result[i] = digits[i] % 10
				aFlag = false
				continue
			}
			result[i] = digits[i]

		}
		return result
	}

}

func PlusOne(digits []int) []int {
	return plusOne(digits)
}
