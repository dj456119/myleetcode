/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-13 00:15:18
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-13 00:37:46
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := -9999999999
	dfs(root, &result)
	return *&result
}

func dfs(root *TreeNode, max *int) int {
	maxSum := root.Val
	m1 := -9999999999
	m2 := -9999999999
	if root.Left != nil {
		m1 = dfs(root.Left, max)
		if m1 > *max {
			*max = m1
		}
	}
	if root.Right != nil {
		m2 = dfs(root.Right, max)
		if m2 > *max {
			*max = m2
		}
	}
	maxSum1 := 0
	if m1 > m2 {
		maxSum = max123(maxSum+m1, maxSum)
		maxSum1 = max123(maxSum, maxSum+m2)
	} else {
		maxSum = max123(maxSum+m2, maxSum)
		maxSum1 = max123(maxSum, maxSum+m1)
	}
	if maxSum > *max {
		*max = maxSum
	}
	if maxSum1 > *max {
		*max = maxSum1
	}
	return maxSum
}

func max123(a, b int) int {
	if a > b {
		return a
	}
	return b
}
