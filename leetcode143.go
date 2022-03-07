/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-05 14:47:57
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-06 01:05:18
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	length := 0
	temp := head
	for temp != nil {
		length++
		temp = temp.Next
	}
	z := 0
	temp = head
	if length%2 == 0 {
		z = length/2 - 1
	} else {
		z = length / 2
	}
	q := 0
	for {
		temp = temp.Next
		q++
		if q == z {
			break
		}
	}
	z3 := rever(temp.Next)
	temp.Next = nil
	result := new(ListNode)
	rtemp := result
	flag := false
	nhead := head
	for {
		if !flag {
			rtemp.Next = nhead
			cc := nhead.Next
			nhead.Next = nil
			nhead = cc
			rtemp = rtemp.Next
			if cc == nil {
				break
			}
		} else {
			rtemp.Next = z3
			cc := z3.Next
			z3.Next = nil
			z3 = cc
			rtemp = rtemp.Next
		}
	}
}

func rever(a *ListNode) *ListNode {
	var result *ListNode
	for {
		if a == nil {
			break
		}
		temp := a.Next
		a.Next = result
		result = a
		a = temp
	}
	return result
}
