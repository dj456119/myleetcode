/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-08 00:46:02
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-08 00:48:15
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	return dfs(root)
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	result = append(result, dfs(root.Left)...)
	result = append(result, dfs(root.Right)...)
	result = append(result, root.Val)
	return result
}
