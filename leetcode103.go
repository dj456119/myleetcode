/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-02 01:29:40
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-03 00:29:25
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	l := 1
	b := false
	for len(queue) != 0 {
		length := l
		l = 0
		result = append(result, []int{})
		if !b {
			for i := length; i > 0; i-- {
				z1 := queue[0]
				result[len(result)-1] = append(result[len(result)-1], z1.Val)
				queue = queue[1:]
				if z1.Left != nil {
					queue = append(queue, z1.Left)
					l++
				}
				if z1.Right != nil {
					queue = append(queue, z1.Right)
					l++
				}
			}
		} else {
			for i := length; i > 0; i-- {
				z1 := queue[len(queue)-1]
				result[len(result)-1] = append(result[len(result)-1], z1.Val)
				queue = queue[:len(queue)-1]
				if z1.Right != nil {
					queue = append([]*TreeNode{z1.Right}, queue...)
					l++
				}
				if z1.Left != nil {
					queue = append([]*TreeNode{z1.Left}, queue...)
					l++
				}
			}
		}
		b = !b

	}
	return result
}
