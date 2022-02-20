/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-20 11:09:06
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-20 11:17:53
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	result := new(ListNode)
	rtemp := result
	for head != nil {
		if head.Val == 0 {
			rtemp.Next = head
			rtemp = rtemp.Next
		} else {
			rtemp.Val = head.Val + rtemp.Val
		}
		head = head.Next
	}
	rtemp = result
	prev := rtemp
	rtemp = rtemp.Next
	for rtemp != nil {
		if rtemp.Val == 0 {
			prev.Next = rtemp.Next
			rtemp = rtemp.Next
		}
		prev = rtemp
		rtemp = rtemp.Next
	}
	return result.Next
}
