package myleetcode

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	r1 := make([]int, len(edges))
	r2 := make([]int, len(edges))

	r1[node1] = 1
	r2[node1] = 1

	n1 := edges[node1]
	n2 := edges[node2]

	l1 := 1
	l2 := 1

	for {
		if n1 == -1 || r1[n1] != 0 {
			break
		}
		l1++
		r1[n1] = l1
		n1 = edges[n1]
	}

	for {
		if n2 == -1 || r2[n2] == 1 {
			break
		}
		l2++
		r2[n2] = l2
		n2 = edges[n2]
	}

	min := 9999999999
	for i := range edges {
		if r1[i] != 0 && r2[i] != 0 {
			m := 0
			if r1[i] < r2[i] {
				m = r1[i] - 1
			} else {
				m = r2[i] - 1
			}
			if m < min {
				min = m
			}
		}
	}
	return min
}
