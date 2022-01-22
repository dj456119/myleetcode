/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-22 17:01:55
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-22 17:10:35
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	result := new(ListNode)
	tempResult := result
	for {
		var min *ListNode
		minIndex := 0
		for i, j := range lists {
			if j == nil {
				continue
			}
			if min == nil || min.Val > j.Val {
				min = j
				minIndex = i
			}
		}
		if min == nil {
			break
		}
		tempResult.Next = min
		tempResult = tempResult.Next
		lists[minIndex] = lists[minIndex].Next
	}

	return result.Next
}
