/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-11-03 00:13:44
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-03 00:44:17
 */
package myleetcode

type ListNode21 struct {
	Val  int
	Next *ListNode21
}

func mergeTwoLists(l1 *ListNode21, l2 *ListNode21) *ListNode21 {
	var newHead *ListNode21
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		newHead = l2
		l2 = l2.Next
	} else {
		newHead = l1
		l1 = l1.Next
	}
	temp := newHead
	for {
		if l1 == nil {
			temp.Next = l2
			return newHead
		}
		if l2 == nil {
			temp.Next = l1
			return newHead
		}
		if l1.Val > l2.Val {
			temp.Next = l2
			temp = l2
			l2 = l2.Next
		} else {
			temp.Next = l1
			temp = l1
			l1 = l1.Next
		}
	}

}

func MergeTwoLists(l1 *ListNode21, l2 *ListNode21) *ListNode21 {
	return mergeTwoLists(l1, l2)
}
