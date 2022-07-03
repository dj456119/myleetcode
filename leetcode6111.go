package myleetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		for j := range result[i] {
			result[i][j] = -1
		}
	}
	a1x := 0
	a1y := 0
	b1x := n - 1
	b1y := a1y
	c1x := b1x
	c1y := m - 1
	d1x := a1x
	d1y := c1y
	node := head
	for {

		if a1x == b1x && a1y == c1y {
			result[a1y][a1x] = node.Val
			return result
		}
		if a1x == b1x {
			// fmt.Println("here")
			for i := a1y; i <= c1y; i++ {
				result[i][a1x] = node.Val
				node = node.Next
				if node == nil {
					return result
				}
			}
		}
		if node == nil {
			break
		}
		for i := a1x; i < b1x; i++ {
			//      fmt.Println(node)
			result[a1y][i] = node.Val
			node = node.Next
			if node == nil {
				return result
			}
		}
		if node == nil {
			break
		}
		for j := b1y; j < c1y; j++ {
			result[j][b1x] = node.Val
			node = node.Next
			if node == nil {
				return result
			}
		}
		if node == nil {
			break
		}
		for i := c1x; i > d1x; i-- {
			result[c1y][i] = node.Val
			node = node.Next
			if node == nil {
				return result
			}
		}
		if node == nil {
			break
		}
		for j := d1y; j > a1y; j-- {
			result[j][d1x] = node.Val
			node = node.Next
			if node == nil {
				return result
			}
		}
		a1x = a1x + 1
		a1y = a1y + 1
		b1x = b1x - 1
		b1y = a1y
		c1x = b1x
		c1y = c1y - 1
		d1x = a1x
		d1y = c1y
	}
	return result
}
