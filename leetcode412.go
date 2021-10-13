/*
 * @Descripttion:给你一个整数 n ，找出从 1 到 n 各个整数的 Fizz Buzz 表示，并用字符串数组 answer（下标从 1 开始）返回结果，其中：
 answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
 answer[i] == "Fizz" 如果 i 是 3 的倍数。
 answer[i] == "Buzz" 如果 i 是 5 的倍数。
 answer[i] == i 如果上述条件全不满足。
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-13 18:15:24
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-13 18:21:57
*/

package myleetcode

import "fmt"

func fizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		mod3 := i % 3
		mod5 := i % 5
		if mod3 == 0 && mod5 == 0 {
			result = append(result, "FizzBuzz")
		} else if mod3 == 0 {
			result = append(result, "Fizz")
		} else if mod5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, fmt.Sprint(i))
		}
	}
	return result
}

func FizzBuzz(n int) []string {
	return fizzBuzz(n)
}
