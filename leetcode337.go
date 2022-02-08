/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-07 00:46:53
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-07 00:55:28
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	a, b := bl(root)
	return max(a, b)
}

//第一个返回值是包含该节点的最大值
//第二个返回值是不包含该节点的最大值
func bl(root *TreeNode) (int, int) {
	includeMax := 0
	noMax := 0
	i1 := 0
	n1 := 0
	i2 := 0
	n2 := 0
	if root.Left != nil {
		i1, n1 = bl(root.Left)
	}
	if root.Right != nil {
		i2, n2 = bl(root.Right)
	}
	//不包含自己
	noMax = max(n1+n2, n1+i2)
	noMax = max(noMax, i1+n2)
	noMax = max(noMax, i1+i2)
	//包含自己
	includeMax = n1 + n2 + root.Val
	return includeMax, noMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
