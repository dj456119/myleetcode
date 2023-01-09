package myleetcode

type Node6285 struct {
	Val    int
	Next   *Node6285
	Random *Node6285
}

func copyRandomList(head *Node6285) *Node6285 {
	z := []*Node6285{}
	t := head
	for t != nil {
		f := new(Node6285)
		f.Val = t.Val
		z = append(z, f)
		t = t.Next
	}
	z = append(z, nil)
	for i := range z {
		if z[i] != nil {
			z[i].Next = z[i+1]
		}
	}
	g := make(map[*Node6285]int)
	t = head
	i := 0
	for t != nil {
		g[t] = i
		i++
		t = t.Next
	}
	g[nil] = i
	t = head
	i = 0
	for t != nil {
		z[i].Random = z[g[t.Random]]
		t = t.Next
		i++
	}
	return z[0]
}
