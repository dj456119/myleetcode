package myleetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	du := make([]int, numCourses)
	for i := range prerequisites {
		du[prerequisites[i][0]]++
	}
	z := []int{}
	for i := range du {
		if du[i] == 0 {
			z = append(z, i)
		}
	}
	result := []int{}
	for len(z) != 0 {
		q := z[0]
		z = z[1:]
		for i := range prerequisites {
			if prerequisites[i][1] == q {
				du[prerequisites[i][0]]--
				if du[prerequisites[i][0]] == 0 {
					z = append(z, prerequisites[i][0])
				}
			}
		}
		result = append(result, q)
	}
	if len(result) != numCourses {
		return []int{}
	}
	return result
}
