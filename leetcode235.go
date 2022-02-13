/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-13 15:46:33
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-13 16:30:55
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return find(root, p, q)
}

func find(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val {
		return p
	}
	if root.Val == q.Val {
		return q
	}
	a1 := find(root.Left, p, q)
	a2 := find(root.Right, p, q)
	if a1 == nil && a2 == nil {
		return nil
	}
	if a1 == nil && a2 != nil {
		return a2
	}
	if a2 == nil && a1 != nil {
		return a1
	}
	return root
}
