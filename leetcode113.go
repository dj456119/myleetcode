/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-08 22:58:10
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-08 23:21:12
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	return dfs(root, targetSum, []int{})
}

func dfs(root *TreeNode, targetSum int, result []int) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		if targetSum-root.Val == 0 {
			result1 := make([]int, len(result))
			copy(result1, result)
			result1 = append(result1, root.Val)
			return [][]int{result1}
		}
		return nil
	}
	rr := [][]int{}
	result = append(result, root.Val)
	a := dfs(root.Left, targetSum-root.Val, result)
	if a != nil {
		rr = append(rr, a...)
	}

	b := dfs(root.Right, targetSum-root.Val, result)
	if b != nil {
		rr = append(rr, b...)
	}
	result = result[:len(result)-1]
	return rr
}
