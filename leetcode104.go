/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 15:09:47
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 15:17:56
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
func maxDepth(root *TreeNode) int {
	max := 1
	nowMaxDepth(root, 1, &max)
	return max
}

func nowMaxDepth(root *TreeNode, nowmax int, max *int) {
	if root != nil {
		nowmax++
		if root.Left == nil && root.Right == nil {
			if nowmax > *max {
				max = &nowmax
			}
		} else {
			nowMaxDepth(root.Left, nowmax, max)
			nowMaxDepth(root.Right, nowmax, max)
		}
	}
}
