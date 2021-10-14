/*
 * @Descripttion:给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

 * @version:
 * @Author: cm.d
 * @Date: 2021-10-14 17:48:21
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-14 18:12:20
*/

package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func merge(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val = root1.Val + root2.Val
	rootLeftIsNull := false
	rootRightIsNUll := false
	if root1.Left == nil {
		root1.Left = root2.Left
		rootLeftIsNull = true
	}
	if root1.Right == nil {
		root1.Right = root2.Right
		rootRightIsNUll = true
	}
	if root1.Left != nil && !rootLeftIsNull {
		merge(root1.Left, root2.Left)
	}
	if root1.Right != nil && !rootRightIsNUll {
		merge(root1.Right, root2.Right)
	}
	return root1
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return merge(root1, root2)
}

func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return mergeTrees(root1, root2)
}
