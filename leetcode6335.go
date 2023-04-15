package myleetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
	parent := make(map[*TreeNode]*TreeNode)
	parent[root] = nil
	level := [][]*TreeNode{[]*TreeNode{}}
	queue := []*TreeNode{root}
	count := 1
	next := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		count--
		level[len(level)-1] = append(level[len(level)-1], node)
		if node.Left != nil {
			parent[node.Left] = node
			queue = append(queue, node.Left)
			next++
		}
		if node.Right != nil {
			parent[node.Right] = node
			queue = append(queue, node.Right)
			next++
		}

		if count == 0 {
			level = append(level, []*TreeNode{})
			count = next
			next = 0
		}
	}
	//  fmt.Println(len(level))

	zmap := make(map[*TreeNode]*TreeNode)
	DeepCopyTree(root, &zmap)
	for i := range level {
		for j := range level[i] {
			// fmt.Print(zmap[level[i][j]], " ")
			sum := 0
			for m := range level[i] {
				if j == m {
					continue
				}
				if parent[level[i][j]] == parent[level[i][m]] {
					continue
				}
				sum = sum + zmap[level[i][m]].Val
			}
			level[i][j].Val = sum
		}
		// fmt.Println(i)
	}
	return root
}

func DeepCopyTree(root *TreeNode, zmap *map[*TreeNode]*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := &TreeNode{Val: root.Val}
	(*zmap)[root] = node
	node.Left = DeepCopyTree(root.Left, zmap)
	node.Right = DeepCopyTree(root.Right, zmap)
	return node
}
