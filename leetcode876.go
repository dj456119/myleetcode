package myleetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	count := 0
	temp := head
	which := 0
	for {
		if temp != nil {
			count++
			temp = temp.Next
		} else {
			break
		}
	}
	which = count / 2
	temp = head
	count = 0
	for {
		if temp != nil {
			if count == which {
				return temp
			}
			count++
			temp = temp.Next
		}
	}
}

func MiddleNode(head *ListNode) *ListNode {
	return middleNode(head)
}
