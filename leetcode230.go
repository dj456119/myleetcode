/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-12 13:08:13
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-12 13:13:35
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	c, _ := findk(root, 0, k)
	return c
}

func findk(root *TreeNode, count int, k int) (int, int) {
	if root == nil {
		return -1, count
	}
	c, d := findk(root.Left, count, k)
	if d == k {
		return c, d
	}
	d++
	if d == k {
		return root.Val, d
	}
	return findk(root.Right, d, k)
}
