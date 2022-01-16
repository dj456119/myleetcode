/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 22:56:52
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 23:21:02
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	left--
	right--
	i := 0
	result := head
	var prev *ListNode
	var leftNode *ListNode
	var rightNode *ListNode
	var leftNode1 *ListNode
	var rightNode1 *ListNode
	for {
		if head == nil {
			break
		}
		if i == left-1 {
			leftNode = head
		}
		if i == right+1 {
			rightNode = head
			break
		}
		if i >= left && i <= right {
			if i == left {
				leftNode1 = head
			}
			if i == right {
				rightNode1 = head
			}
			next := head.Next
			head.Next = prev
			head = next
			prev = head
		} else {
			head = head.Next
		}
		i++
	}
	if leftNode != nil {
		leftNode.Next = rightNode1
	}
	if rightNode != nil {
		rightNode.Next = leftNode1
	}
	return result
}

// func reverse(head *ListNode) *ListNode {
// 	var tail *ListNode
// 	for {
// 		if head == nil {
// 			break
// 		}
// 		next := head.Next
// 		head.Next = tail
// 		tail = head
// 		if next == nil {
// 			return head
// 		}
// 		head = next
// 	}
// 	return nil
// }
