/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-30 09:16:24
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-30 09:39:53
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	head1 := new(ListNode)
	head1.Next = head
	head1 = head
	fast, slow, pos := head, head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	for pos != slow {
		pos = pos.Next
		slow = slow.Next
	}
	return pos
}
