/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 17:05:17
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 17:06:45
 */
package myleetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	v := []int{}
	for {
		if head == nil {
			break
		}
		v = append(v, head.Val)
		head = head.Next
	}

	i := 0
	j := len(v) - 1
	for {
		if v[i] != v[j] {
			return false
		}
		i++
		j--
		if i >= j {
			return true
		}
	}
	return true
}
