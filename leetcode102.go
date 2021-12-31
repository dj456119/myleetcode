/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-31 19:51:05
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-31 23:54:46
 */
package myleetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type LevelTreeNode struct {
	level int
	TreeNode
}

func NewLevelTreeNode(root *TreeNode, level int) *LevelTreeNode {
	l := LevelTreeNode{}
	l.Val = root.Val
	l.Left = root.Left
	l.Right = root.Right
	l.level = level
	return &l
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	temp := []*LevelTreeNode{}
	result := [][]int{}
	temp = append(temp, NewLevelTreeNode(root, 0))

	for len(temp) != 0 {
		t := temp[0]
		if len(result) == t.level {
			result = append(result, []int{t.Val})
		} else {
			result[t.level] = append(result[t.level], t.Val)
		}
		if t.Left != nil {
			temp = append(temp, NewLevelTreeNode(t.Left, t.level+1))
		}
		if t.Right != nil {
			temp = append(temp, NewLevelTreeNode(t.Right, t.level+1))
		}
		temp = temp[1:len(temp)]
	}
	return result
}
