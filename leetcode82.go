/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-22 17:22:18
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-22 17:30:11
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	nh := new(ListNode)
	nh.Next = head
	prev := nh
	for {
		if head == nil {
			break
		}
		if head.Next == nil {
			break
		}
		if head.Next.Val == head.Val {
			for {
				if head.Next == nil {
					break
				}
				if head.Val != head.Next.Val {
					break
				}
				head.Next = head.Next.Next
			}
			prev.Next = head.Next
			head = head.Next
		} else {
			prev = head
			head = head.Next
		}
	}
	return nh.Next
}
