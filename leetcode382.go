/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-16 00:09:31
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-16 00:09:31
 */
package myleetcode

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (s *Solution) GetRandom() (ans int) {
	for node, i := s.head, 1; node != nil; node = node.Next {
		if rand.Intn(i) == 0 {
			ans = node.Val
		}
		i++
	}
	return
}
