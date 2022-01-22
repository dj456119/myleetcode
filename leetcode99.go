/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-23 00:04:09
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-23 00:25:55
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	r := dfs(root, []int{})
	i := 0
	j := 1
	a1 := []int{}
	for {
		if j == len(r) {
			break
		}
		if r[i] > r[j] {
			if len(a1) == 1 {
				a1 = append(a1, r[j])
			} else {
				a1 = append(a1, r[i])
			}

		}
		if len(a1) == 2 {
			break
		}
		i++
		j++
	}
	if len(a1) == 1 {
		for i, j := range r {
			if j == a1[0] {
				a1 = append(a1, r[i+1])
				break
			}
		}
	}

	recovery(root, a1)
}

func dfs(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}
	r := dfs(root.Left, result)
	r = append(r, root.Val)
	r = dfs(root.Right, r)
	return r
}

func recovery(root *TreeNode, a []int) {
	if root == nil {
		return
	}
	recovery(root.Left, a)

	if root.Val == a[0] {
		root.Val = a[1]
	} else if root.Val == a[1] {
		root.Val = a[0]
	}

	recovery(root.Right, a)
}
