package myleetcode

type Node133 struct {
	Val       int
	Neighbors []*Node133
}

func cloneGraph(node *Node133) *Node {
	temp := make(map[int]map[int]*Node133)
	next := node
	var prev *Node133
	for next != nil {
		var nNode *Node133
		if z, ok := temp[next.Val]; ok {
			nNode = z[next.Val]
		} else {
			nNode = new(Node133)
			temp[next.Val] = make(map[int]*Node133)
			temp[next.Val][next.Val] = nNode
			nNode.Val = next.Val
			nNode.Neighbors = []*Node133{}
		}
		list := next.Neighbors
		for _, h := range list {
			if z, ok := temp[h.Val]; ok {
				c := temp[h.Val][h.Val]

			}
		}
	}
}
