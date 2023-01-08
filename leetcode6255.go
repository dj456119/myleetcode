package myleetcode

func minScore(n int, roads [][]int) int {
	min := 99999999999
	z := make(map[int][][]int)
	for i := range roads {
		if r, ok := z[roads[i][0]]; ok {
			r = append(r, []int{roads[i][1], roads[i][2]})
			z[roads[i][0]] = r
		} else {
			r := [][]int{}
			r = append(r, []int{roads[i][1], roads[i][2]})
			z[roads[i][0]] = r
		}
	}
	for i := range roads {
		if r, ok := z[roads[i][1]]; ok {
			r = append(r, []int{roads[i][0], roads[i][2]})
			z[roads[i][1]] = r
		} else {
			r := [][]int{}
			r = append(r, []int{roads[i][0], roads[i][2]})
			z[roads[i][1]] = r
		}
	}
	queue := [][]int{}
	queue = append(queue, z[1]...)
	f := make(map[int]bool)
	f[1] = true
	for len(queue) != 0 {
		if queue[0][1] < min {
			min = queue[0][1]
		}
		if _, ok := f[queue[0][0]]; !ok {
			queue = append(queue, z[queue[0][0]]...)
			f[queue[0][0]] = true
		}
		queue = queue[1:]
	}
	return min
}
