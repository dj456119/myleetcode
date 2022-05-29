package myleetcode

import "sort"

func largestWordCount(messages []string, senders []string) string {
	cmap := make(map[string]int)
	for i := range messages {
		cmap[senders[i]] = cmap[senders[i]] + getWCount(messages[i])
	}
	max := 0
	for _, v := range cmap {
		if v > max {
			max = v
		}
	}
	z := []string{}
	for k, v := range cmap {
		if v == max {
			z = append(z, k)
		}
	}
	sort.Strings(z)
	return z[len(z)-1]
}

func getWCount(s string) int {
	sum := 0
	for i := range s {
		if s[i] == ' ' {
			sum++
		}
	}
	return sum + 1
}
