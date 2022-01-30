/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-22 01:37:23
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-30 10:18:38
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	head1 := new(ListNode)
	head1.Next = head
	temp := head1
	step := 1
	length := 0
	for temp.Next != nil {
		temp = temp.Next
		length++
	}

	for step < length {
		l1 := new(ListNode)
		l2 := new(ListNode)
		tl1 := l1
		tl2 := l2
		temp = head1
		for {
			for i := 0; i < step; i++ {
				if temp.Next == nil {
					break
				}
				l1.Next = temp.Next
				l1 = l1.Next
				temp = temp.Next
				if temp.Next == nil {
					break
				}
				l2.Next = temp.Next
				l2 = l2.Next
				temp = temp.Next
			}
		}
	}
}

func merge(l1, l2 *ListNode, step int) (*ListNode, *ListNode) {
	if l2 == nil {
		return l1, l1
	}
	result := new(ListNode)
	t := result
	for {
		if l1 == nil {
			t.Next = l2
			break
		}
		if l2 == nil {
			t.Next = l1
			break
		}
		if l1.Val > l2.Val {
			t.Next = l2
		} else {
			t.Next = l1
		}
		t = t.Next
	}
	t = result
	for t.Next != nil {
		t = t.Next
	}
	return result.Next, t
}
