/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 10:34:08
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 19:48:00
 */
package myleetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	head := new(ListNode)
	result := head
	for {
		if l1 == nil || l2 == nil {
			break
		}
		sum := l1.Val + l2.Val + flag
		ln := new(ListNode)
		ln.Val = sum % 10
		flag = sum / 10
		head.Next = ln
		l1 = l1.Next
		l2 = l2.Next
		head = head.Next
	}
	if l2 != nil {
		l1 = l2
	}
	for l1 != nil {
		sum := l1.Val + flag
		ln := new(ListNode)
		ln.Val = sum % 10
		flag = sum / 10
		head.Next = ln
		l1 = l1.Next
		head = head.Next
	}
	if flag != 0 {
		ln := new(ListNode)
		ln.Val = 1
		head.Next = ln
	}
	return result.Next
}
