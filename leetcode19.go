package myleetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n == 0 {
		return head
	}
	var temp, temp2 *ListNode
	temp = head
	temp2 = head
	count := 0
	count2 := 0
	for {
		if temp != nil {
			temp = temp.Next
			count++
			if temp == nil {
				if n == count {
					return head.Next
				}
				temp2.Next = temp2.Next.Next
				break
			}
			if count >= n+1 {
				temp2 = temp2.Next
				count2++
			}
		}
	}
	return head
}

func InitNode() *ListNode {
	nums := []int{1, 2}
	var temp, next *ListNode

	for i := len(nums) - 1; i >= 0; i-- {
		temp = new(ListNode)
		temp.Val = nums[i]
		temp.Next = next
		next = temp
	}
	return temp
}

func PrintNode(head *ListNode) {
	if head == nil {
		fmt.Println("[]")
		return
	}
	temp := head
	for {
		fmt.Print(temp.Val)
		fmt.Print("=>")
		temp = temp.Next
		if temp == nil {
			break
		}
	}
	fmt.Println("EOF")
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEnd(head, n)
}
