/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 10:46:01
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 10:56:10
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
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		r1 := inorderTraversal(root.Left)
		result = append(result, r1...)
		result = append(result, root.Val)
		r2 := inorderTraversal(root.Right)
		result = append(result, r2...)
	}
	return result
}
