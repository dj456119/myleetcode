/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-01 15:06:55
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-01 15:14:34
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	result := zx(root, []int{})
	for i := len(result) - 1; i >= 0; i++ {
		if i != len(result)-1 {
			result[i] = result[i] + result[i-1]
		}
	}
	zx1(root, result, 0)
	return root
}

func zx(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}
	result = zx(root.Left, result)
	result = append(result, root.Val)
	return zx(root.Right, result)
}

func zx1(root *TreeNode, result []int, pos int) int {
	if root == nil {
		return pos
	}
	pos = zx1(root.Left, result, pos)
	root.Val = result[pos]
	return zx1(root.Right, result, pos+1)
}
