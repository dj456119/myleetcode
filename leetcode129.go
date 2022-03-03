/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-04 00:06:12
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-04 00:37:36
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, value int) int {
	value = value*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return value
	}
	l := 0
	if root.Left != nil {
		l = dfs(root.Left, value)
	}
	r := 0
	if root.Right != nil {
		r = dfs(root.Right, value)
	}
	z := l + r
	return z
}
