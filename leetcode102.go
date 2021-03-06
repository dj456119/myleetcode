/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 19:51:05
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-02 01:27:04
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	l := 1
	for len(queue) != 0 {
		length := l
		l = 0
		result = append(result, []int{})
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
	}
	return result
}
