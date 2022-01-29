/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-25 00:49:57
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-25 23:49:50
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	result := dsf(root, []*TreeNode{})
	z := result[0]
	for i, j := range result {
		if i != 0 {
			z.Right = j
			z = j
		}
	}
}

func dsf(root *TreeNode, result []*TreeNode) []*TreeNode {
	if root == nil {
		return result
	}
	result = append(result, root)
	l1 := dsf(root.Left, result)
	return dsf(root.Right, l1)
}
