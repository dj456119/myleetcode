package main

func vowelStrings(words []string, left int, right int) int {
	z := []byte{'a', 'e', 'i', 'o', 'u'}
    result := 0
	for i := range words {
		if i >= left && i <= right {
			q := 0
			for j := range z {
				if z[j] == words[i][0]  {
					q++
				}
                if  z[j] == words[i][len(words[i]) - 1] {
                    q++
                }
			}
			if q == 2 {
				result++
			}
		}
	}
	return result
}