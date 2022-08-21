package myleetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNodeData6154 struct {
	Self, Left, Right, Father *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	bmap := make(map[int]bool)
	z := make(map[*TreeNode]*TreeNodeData6154)
	rootTnd := new(TreeNodeData6154)
	rootTnd.Father = nil
	rootTnd.Self = root

	queue := []*TreeNodeData6154{rootTnd}
	for len(queue) != 0 {
		tnd := queue[0]
		queue = queue[1:]
		z[tnd.Self] = tnd
		rootTnd.Left = root.Left
		rootTnd.Right = root.Right
	}
}
