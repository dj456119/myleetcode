/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 10:38:58
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 10:45:50
 */
package myleetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		b1 := isSameTree(p.Left, q.Left)
		b2 := isSameTree(p.Right, q.Right)
		return b1 && b2
	}
	return false
}
