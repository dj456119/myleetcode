/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-08 23:14:58
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-08 23:18:35
 */
package myleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	ra := []int{}
	for {
		if head == nil {
			break
		}
		ra = append(ra, head.Val)
		head = head.Next
	}
	i := 0
	j := len(ra) - 1
	max := 0
	start := false
	for {
		sum := ra[i] + ra[j]
		if !start || sum > max {
			max = sum
			start = true
		}
		i++
		j--
		if i >= j {
			break
		}
	}
	return max
}
