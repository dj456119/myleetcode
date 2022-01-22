/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-22 01:37:23
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-22 02:00:56
 */
package myleetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	result := new(ListNode)
	temp := result
	z := new(ListNode)
	z.Next = head
	for {
		min := findMin(z)
		fmt.Println(min)

		temp.Next = min
		if z.Next == nil {
			break
		}

		temp = temp.Next
	}
	return result.Next
}

func findMin(z *ListNode) *ListNode {
	var minPrev *ListNode = z
	var min *ListNode = z.Next
	var head *ListNode = z.Next
	var prev *ListNode = z
	for {
		if head == nil {
			break
		}
		if min == nil || head.Val <= min.Val {
			minPrev = prev
			min = head
		}
		prev = head
		head = head.Next
	}
	if minPrev != nil {
		minPrev.Next = min.Next
	}
	return min
}
