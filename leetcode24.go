/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-08 17:00:53
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-08 17:05:21
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	result := head.Next
	var prev *ListNode
	for {
		first := head
		if first == nil {
			return result
		}
		second := head.Next
		if second == nil {
			return result
		}
		head = second.Next
		if prev == nil {
			prev = first
		} else {
			prev.Next = second
			prev = first
		}
		second.Next = first
		first.Next = head
	}
}
