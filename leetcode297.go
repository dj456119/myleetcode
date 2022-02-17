/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-18 02:24:44
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-18 02:25:50
 */
package myleetcode

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	re := qianxu(root, []*TreeNode{})
	result := ""
	for i := range re {
		if re[i] == nil {
			result = fmt.Sprintf("%s,nil", result)
		} else {
			result = fmt.Sprintf("%s,%d", result, re[i].Val)
		}
	}
	return result[1:]
}

func qianxu(root *TreeNode, result []*TreeNode) []*TreeNode {
	if root == nil {
		return append(result, nil)
	}
	result = append(result, root)
	result = qianxu(root.Left, result)
	return qianxu(root.Right, result)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	reString := strings.Split(data, ",")
	if len(reString) == 0 {
		return nil
	}
	te := []*TreeNode{}
	for i := range reString {
		if reString[i] == "nil" {
			te = append(te, nil)
		} else {
			tr := new(TreeNode)
			tz, _ := strconv.ParseInt(reString[i], 10, 64)
			tr.Val = int(tz)
			te = append(te, tr)
		}
	}
	root := te[0]
	if root == nil {
		return nil
	}
	buildTree(root, 0, te)
	return te[0]
}

func buildTree(root *TreeNode, pos int, te []*TreeNode) int {
	if pos == len(te) {
		return pos
	}
	if te[pos+1] == nil {
		root.Left = nil
		pos++
	} else {
		root.Left = te[pos+1]
		pos = buildTree(root.Left, pos+1, te)
	}
	if te[pos+1] == nil {
		root.Right = nil
		pos++
	} else {
		root.Right = te[pos+1]
		pos = buildTree(root.Right, pos+1, te)
	}
	return pos
}
