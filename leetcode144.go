/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-22 16:44:29
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-22 16:47:57
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return dfs(root, []int{})
}

func dfs(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	r1 := dfs(root.Left, result)
	r2 := dfs(root.Right, r1)
	return r2
}
