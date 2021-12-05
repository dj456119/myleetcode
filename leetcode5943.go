/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-05 14:20:14
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-05 14:20:14
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	if head.Next.Next == nil {
		head.Next = nil
		return head
	}
	headTmp := head
	count := 0
	for {
		if headTmp != nil {
			count++
			headTmp = headTmp.Next
		} else {
			break
		}
	}
	pos := 0
	if count%2 == 0 {
		pos = count / 2
	} else {
		pos = count / 2
	}

	headTmp = head
	for {
		if pos == 1 {
			headTmp.Next = headTmp.Next.Next
			break
		}
		headTmp = headTmp.Next
		pos--
	}
	return head
}
