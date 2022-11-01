package myleetcode

import "sort"

type creator struct {
	Sum int
	ids []int
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	max := 0
	maxC := []string{}
	z := make(map[string]creator)
	for i := range creators {
		if c, ok := z[creators[i]]; ok {
			c.Sum = c.Sum + views[i]
			c.ids = append(c.ids, i)
			if c.Sum > max {
				max = c.Sum
			}
			z[creators[i]] = c
		} else {
			c = creator{
				Sum: views[i],
				ids: []int{i},
			}
			z[creators[i]] = c
			if views[i] > max {
				max = views[i]
			}
		}
	}
	//   fmt.Println(z)
	for k, c := range z {
		if c.Sum == max {
			maxC = append(maxC, k)
		}
	}
	result := [][]string{}
	for j := range maxC {
		iids := z[maxC[j]].ids
		maxii := 0
		f := []string{}
		for i := range iids {
			if views[iids[i]] > maxii {
				maxii = views[iids[i]]
			}
		}
		for i := range iids {
			if views[iids[i]] == maxii {
				f = append(f, ids[iids[i]])
			}
		}
		sort.Slice(f, func(i, j int) bool { return f[i] < f[j] })
		result = append(result, []string{maxC[j], f[0]})
	}

	return result
}
