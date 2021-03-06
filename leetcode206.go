/*
 * @Descripttion:给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 * @version:
 * @Author: cm.d
 * @Date: 2021-11-03 20:26:20
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-17 01:02:40
 */

package myleetcode

type ListNode206 struct {
	Val  int
	Next *ListNode206
}

func reverseList1(head *ListNode206) *ListNode206 {
	var prev *ListNode206

	for {
		if head == nil {
			return prev
		}
		if head.Next == nil {
			head.Next = prev
			return head
		}
		next := head.Next
		head.Next = prev
		prev = next
		head, next.Next = next.Next, head
	}
}

func reverseList(head *ListNode) *ListNode {
	var tail *ListNode
	for {
		if head == nil {
			break
		}
		next := head.Next
		head.Next = tail
		tail = head
		head = next
	}
	return tail
}

func ReverseList(head *ListNode206) *ListNode206 {
	return reverseList1(head)
}
