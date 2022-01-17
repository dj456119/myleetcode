/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-18 00:15:30
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-18 00:38:28
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	r := dfs(root, []int{})
	for i, j := range r {
		if i != 0 {
			if j <= r[i-1] {
				return false
			}
		}
	}
	return true
}

func dfs(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}
	r1 := dfs(root.Left, result)
	r1 = append(r1, root.Val)
	r1 = dfs(root.Right, r1)
	return r1
}
