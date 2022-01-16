/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-13 19:51:40
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-13 23:51:51
 */
package myleetcode

import (
	"fmt"
	"strconv"
)

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	b := digui([]byte(num), 0, 0, 1, 1, 2, 2)
	fmt.Println("doujieshule", b)
	return b
}
func digui(nums []byte, oneStart, oneEnd, twoStart, twoEnd, sumStart, sumEnd int) bool {
	if len(nums[oneStart:oneEnd+1]) > len(nums[sumStart:]) {
		return false
	}
	if len(nums[twoStart:twoEnd+1]) > len(nums[sumStart:]) {
		return false
	}
	if sumEnd >= len(nums) {
		return false
	}
	one, _ := strconv.ParseInt(string(nums[oneStart:oneEnd+1]), 10, 32)
	two, _ := strconv.ParseInt(string(nums[twoStart:twoEnd+1]), 10, 32)
	sum, _ := strconv.ParseInt(string(nums[sumStart:sumEnd+1]), 10, 32)
	//if one == 12 && two == 0 {
	fmt.Println(one, two, sum)
	//  }
	if one+two == sum && (nums[oneStart] != '0' || len(nums[oneStart:oneEnd+1]) == 1) && (nums[twoStart] != '0' || len(nums[twoStart:twoEnd+1]) == 1) && (nums[sumStart] != '0' || len(nums[sumStart:sumEnd+1]) == 1) {
		fmt.Println(one, two, sum)
		if sumEnd == len(nums)-1 {
			return true
		}
		if digui(nums, twoStart, twoEnd, sumStart, sumEnd, sumEnd+1, sumEnd+1) {
			return true
		}
	}

	if digui(nums, oneStart, oneEnd+1, twoStart+1, twoEnd+1, sumStart+1, sumEnd+1) {
		return true
	}

	if digui(nums, oneStart, oneEnd, twoStart, twoEnd+1, sumStart+1, sumEnd+1) {
		return true
	}

	if digui(nums, oneStart, oneEnd, twoStart, twoEnd, sumStart, sumEnd+1) {
		return true
	}

	return false
}
