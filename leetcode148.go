/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-22 01:37:23
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-31 01:41:31
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
	prev := head1
	for step < length {
		temp = head1.Next
		for temp != nil {
			l1 := new(ListNode)
			l2 := new(ListNode)
			tl1 := l1
			tl2 := l2
			for i := 0; i < step && temp != nil; i++ {
				tl1.Next = temp
				temp = temp.Next
				tl1 = tl1.Next
			}
			tl1.Next = nil
			for i := 0; i < step && temp != nil; i++ {
				tl2.Next = temp
				temp = temp.Next
				tl2 = tl2.Next
			}
			tl2.Next = nil
			nl, ne := merge(l1.Next, l2.Next)
			prev.Next = nl
			prev = ne
		}
		step = step * 2
		prev = head1
	}

	return head1.Next
}

func merge(l1, l2 *ListNode) (*ListNode, *ListNode) {
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
			l2 = l2.Next
		} else {
			t.Next = l1
			l1 = l1.Next
		}
		t = t.Next
	}
	t = result
	for t.Next != nil {
		t = t.Next
	}
	return result.Next, t
}
