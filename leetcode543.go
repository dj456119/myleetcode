/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 17:55:37
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 17:59:07
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package myleetcode

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	diameter(root, &result)
	return result
}

func diameter(root *TreeNode, max *int) int {
	if root != nil {
		lmax := diameter(root.Left, max)
		rmax := diameter(root.Right, max)
		nm := lmax + rmax + 1
		if nm > *max {
			*max = nm
		}
		return nm
	}
	return 0
}
