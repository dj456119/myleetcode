/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 16:28:50
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 16:31:02
 */
package myleetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ta := headA
	tb := headB
	tam := make(map[*ListNode]bool)
	for {
		if ta != nil {
			tam[ta] = true
		} else {
			break
		}
		ta = ta.Next
	}
	for {
		if tb != nil {
			if _, ok := tam[tb]; ok {
				return tb
			}
		} else {
			break
		}
		tb = tb.Next
	}
	return nil
}
