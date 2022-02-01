/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-01 14:05:58
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-01 15:03:24
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return dfs(root, p, q)
}

func dfs(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	a1 := dfs(root.Left, p, q)
	a2 := dfs(root.Right, p, q)
	if a1 != nil && a2 != nil {
		return root
	}
	if a1 == nil && a2 != nil {
		return a2
	} else {
		return a1
	}
}
