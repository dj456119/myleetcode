/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-03 00:33:39
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-03 00:38:05
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	l := 1
	ltemp := l
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		l = 0
		result = append(result, []int{})
		for i := ltemp; i > 0; i-- {
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
		ltemp = l
	}
	i := 0
	j := len(result) - 1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return result
}
