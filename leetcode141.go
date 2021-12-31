/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 15:20:43
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 16:05:42
 */
package myleetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	t1 := head
	t2 := head
	for {
		if t1 == nil {
			return false
		}
		if t1.Next == nil {
			return false
		}
		if t1.Next.Next == nil {
			return false
		}
		t1 = t1.Next.Next
		t2 = t2.Next
		if t1 == t2 {
			return true
		}
	}
}
