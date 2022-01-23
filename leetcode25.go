/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-23 23:10:20
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-24 00:06:24
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	nh := new(ListNode)
	nh.Next = head
	temp := head
	z := k
	result := new(ListNode)
	rr := result
	tempNew := new(ListNode)
	ttn := tempNew
	for {
		if temp == nil && z != 0 {
			rr.Next = tempNew.Next
			break
		}
		if z == 0 {
			r, ns := reverseList(tempNew.Next)
			rr.Next = r
			rr = ns
			z = k
			tempNew.Next = nil
			ttn = tempNew
			if temp == nil {
				break
			}
		}
		next := temp.Next
		temp.Next = nil
		ttn.Next = temp
		ttn = ttn.Next
		temp = next
		z--
	}
	return result.Next
}

func reverseList(head *ListNode) (*ListNode, *ListNode) {
	temp := head
	result := new(ListNode)
	for {
		if temp == nil {
			break
		}
		next := temp.Next
		temp.Next = result.Next
		result.Next = temp
		temp = next
	}
	return result.Next, head
}
