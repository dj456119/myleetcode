package main

type TreeNode struct {
    Val int
    Left *TreeNode
	Right *TreeNode
}
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	ll := []int64{}
	q := []*TreeNode{root}
	nowCount := 1
	nextCount := 0
	var sum int64 = 0
	for len(q) != 0 {
		z := q[0]
		q = q[1:]
        sum = sum + int64(z.Val)
		nowCount = nowCount - 1
		if z.Left != nil {
			nextCount++
			q = append(q, z.Left) 
		}
		if z.Right != nil {
			nextCount++
			q = append(q, z.Right)
		}
		if nowCount == 0 {
			ll = append(ll, sum)
			sum = 0
			nowCount = nextCount
			nextCount = 0
		}
	}
	if len(ll) < k {
		return -1
	}
	sort.Slice(ll, func(x, y int) bool {
		return ll[x] > ll[y]
	} )
	return ll[k-1]
}