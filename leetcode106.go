/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-12 10:28:52
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-12 13:14:49
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, postorder)
}

func build(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = postorder[len(postorder)-1]
	postorder = postorder[:len(postorder)-1]
	for i := range inorder {
		if inorder[i] == root.Val {
			rightLength := len(inorder) - i - 1
			rightIndex := len(postorder) - rightLength
			left := build(inorder[:i], postorder[:rightIndex])
			root.Left = left
			if i != len(inorder)-1 {
				right := build(inorder[i+1:], postorder[rightIndex:])
				root.Right = right
			}
			break
		}
	}
	return root
}
