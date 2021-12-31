/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 14:52:32
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 15:07:33
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
func isSymmetric(root *TreeNode) bool {
	return compareRoot(root.Left, root.Right)
}

func compareRoot(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	if r1.Val == r2.Val {
		b1 := compareRoot(r1.Left, r2.Right)
		b2 := compareRoot(r1.Right, r2.Left)
		if b1 && b2 {
			return true
		} else {
			return false
		}
	}
	return false
}
