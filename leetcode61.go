/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-18 00:05:47
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-18 00:10:21
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	length := 0
	temp1 := head
	var last *ListNode
	for temp1 != nil {
		length++
		last = temp1
		temp1 = temp1.Next
	}
	k1 := length - k
	i := 0
	temp1 = head
	for i != k1 {
		temp1 = temp1.Next
		i++
	}
	next := temp1.Next
	last.Next = head
	temp1.Next = nil
	return next
}
