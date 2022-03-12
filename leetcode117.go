/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-08 23:23:41
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-12 10:07:01
 */
package myleetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	dg(root)
	return root
}

func dg(root *Node) {
	first := root
	for {
		var nextLineNext *Node
		now := first
		first = nil
		for now != nil {
			if nextLineNext != nil {
				if now.Left != nil {
					nextLineNext.Next = now.Left
					nextLineNext = now.Left
				}
				if now.Right != nil {
					nextLineNext.Next = now.Right
					nextLineNext = now.Right
				}
			} else {
				if now.Left != nil {
					nextLineNext = now.Left
					if first == nil {
						first = nextLineNext
					}
					if now.Right != nil {
						nextLineNext.Next = now.Right
						nextLineNext = now.Right
					}

				} else {
					if now.Right != nil {
						nextLineNext = now.Right
						if first == nil {
							first = nextLineNext
						}
					}
				}
			}
			now = now.Next
		}
		if nextLineNext == nil {
			break
		}
	}
}
