/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-03-08 00:21:13
 * @LastEditors: cm.d
 * @LastEditTime: 2022-03-08 00:33:07
 */
package myleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums)
}

func build(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := (len(nums) - 1) / 2
	tn := new(TreeNode)
	tn.Val = nums[mid]
	if mid != 0 {
		tn.Left = build(nums[0:mid])
	}
	if mid != len(nums)-1 {
		tn.Right = build(nums[mid+1:])
	}
	return tn
}
