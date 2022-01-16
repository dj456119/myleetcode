/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 18:21:52
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 18:48:39
 */
package myleetcode

import "fmt"

func myAtoi(s string) int {
	clean := []byte{}
	flag := false
	start := false
	space := false
	for _, j := range []byte(s) {
		if j == ' ' {
			space = true
			continue
		}
		if j == '+' {
			if start && space {
				return 0
			}
			start = true
			continue
		}
		if j == '-' {
			if start || space {
				return 0
			}
			start = true
			flag = true
		}

		if j >= '0' && j <= '9' {
			start = true
			clean = append(clean, j)
		} else {
			if start == false {
				return 0
			} else {
				break
			}
		}

	}
	fmt.Println(string(clean))
	index := 0
	for z, j := range clean {
		if j != '0' {
			break
		}
		index = z
	}
	if index != 0 {
		clean = clean[index+1:]
	}

	fmt.Println(string(clean))
	if len(clean) > 10 {
		if flag {
			return -2147483648
		} else {
			return 2147483647
		}
	}
	sum := 0
	for i, j := range clean {
		if i == 0 {
			sum = int(j - '0')
		} else {
			sum = sum*10 + int(j-'0')
		}
	}
	if flag {
		if sum < -2147483648 {
			sum = -2147483648
		} else {
			sum = -sum
		}
	} else {
		if sum > 2147483647 {
			return 2147483647
		}
	}
	return sum
}
