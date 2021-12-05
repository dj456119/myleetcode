/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-12-05 14:21:44
 * @LastEditors: cm.d
 * @LastEditTime: 2021-12-05 14:22:47
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	for {
		if root.Val == startValue || root.Val == destValue {
			break
		}
		if root.Left != nil && root.Right != nil {
			break
		}
		if root.Left == nil {
			root = root.Right
		}
		if root.Right == nil {
			root = root.Left
		}
	}
	result, _ := findRoot(root, startValue, destValue)
	return result
}

func findLeft(root *TreeNode, value int) ([]byte, bool) {

	if root.Val == value {
		return nil, true
	}
	if root.Left != nil {
		result, b := findLeft(root.Left, value)
		if b {
			return append([]byte{'U'}, result...), true
		}
	}
	if root.Right != nil {

		result, b := findLeft(root.Right, value)
		if b {
			return append([]byte{'U'}, result...), true
		}
	}
	return nil, false
}

func findRight(root *TreeNode, value int) ([]byte, bool) {
	if root.Val == value {
		return nil, true
	}
	if root.Left != nil {
		result, b := findRight(root.Left, value)
		if b {
			return append([]byte{'L'}, result...), true
		}
	}
	if root.Right != nil {
		result, b := findRight(root.Right, value)
		if b {
			return append([]byte{'R'}, result...), true
		}
	}
	return nil, false
}

func findRoot(root *TreeNode, v1, v2 int) (string, bool) {
	b1, b2 := false, false
	var resultLeft, resultRight []byte
	resultString, rr1, rr2 := "", "", ""
	if v1 != root.Val {
		resultLeft, b1 = findLeft(root, v1)
		resultString = resultString + string(resultLeft)
	} else {
		b1 = true
	}
	if v2 != root.Val {
		resultRight, b2 = findRight(root, v2)
		resultString = resultString + string(resultRight)
	} else {
		b2 = true
	}

	if !b1 || !b2 {
		return "", false
	}
	b1, b2 = false, false
	if root.Left != nil {
		rr1, b1 = findRoot(root.Left, v1, v2)
	}
	if root.Right != nil {
		rr2, b2 = findRoot(root.Right, v1, v2)
	}

	if b1 {
		return rr1, true
	}
	if b2 {
		return rr2, true
	}
	return resultString, true
}
