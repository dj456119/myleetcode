/**
给定一个整数，编写算法转换为十六进制数
*/
package myleetcode

import "fmt"

func toHex(num int) string {
	dataMap := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	x := num
	var result string
	if x < 0 {
		x = 0x100000000 + x
	}

	for {
		data := x % 16
		result = fmt.Sprintf("%s%s", dataMap[data], result)
		x = x / 16
		if x == 0 {
			break
		}
	}
	return result
}

func ToHex(num int) string {
	return toHex(num)
}
