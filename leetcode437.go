/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-06 22:24:55
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-06 23:48:18
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	qzh := make(map[int]int)
	qzh[0] = 1
	return dfs(root, 0, targetSum, qzh)
}

func dfs(root *TreeNode, sum, targetSum int, qzh map[int]int) int {
	if root == nil {
		return 0
	}
	r1 := 0
	r1 += qzh[sum+root.Val-targetSum]
	qzh[sum+root.Val]++
	r1 += dfs(root.Left, sum+root.Val, targetSum, qzh)
	r1 += dfs(root.Right, sum+root.Val, targetSum, qzh)
	qzh[sum+root.Val]--
	return r1
}
