package myleetcode

import "sort"

func similarPairs(words []string) int {
	z := make(map[string][]int)
	for i := range words {
		tmp := make(map[byte]bool)
		for j := range words[i] {
			tmp[words[i][j]] = true
		}
		q := []byte{}
		for k, _ := range tmp {
			q = append(q, k)
		}
		sort.Slice(q, func(i, j int) bool { return q[i] < q[j] })
		key := string(q)
		if e, ok := z[key]; ok {
			e = append(e, i)
			z[key] = e
		} else {
			e = []int{}
			e = append(e, i)
			z[key] = e
		}
	}

	result := 0
	for _, v := range z {
		if len(v) == 1 {
			continue
		}
		result = result + (len(v) * (len(v) - 1) / 2)
	}
	return result
}
