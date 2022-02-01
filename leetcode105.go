/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-01 18:33:27
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-01 18:46:33
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder)
}

func build(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		tr := new(TreeNode)
		tr.Val = preorder[0]
		return tr
	}

	index := 0
	for i, j := range inorder {
		if j == preorder[0] {
			index = i
			break
		}
	}
	tr := new(TreeNode)
	tr.Val = preorder[0]
	leftLength := index
	left := build(preorder[1:leftLength+1], inorder[:index])
	if index != len(inorder)-1 {
		right := build(preorder[leftLength+1:], inorder[index+1:])
		tr.Right = right
	}

	tr.Left = left

	return tr
}
