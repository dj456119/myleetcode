/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-18 00:44:32
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-18 01:02:42
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	small := new(ListNode)
	small1 := small
	large := new(ListNode)
	large1 := large
	i := 0
	for {
		i++
		if head == nil {
			break
		}
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	if i == 1 {
		return head
	}
	small.Next = large1.Next
	return small1.Next
}
