/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 11:28:30
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 11:33:05
 */
package myleetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	result := head
	tempMap := make(map[int]bool)
	var prev *ListNode
	for {
		if head != nil {
			if _, ok := tempMap[head.Val]; ok {
				if prev != nil {
					prev.Next = head.Next
				}
			} else {
				tempMap[head.Val] = true
				prev = head
			}
			head = head.Next
		} else {
			break
		}
	}
	return result
}
