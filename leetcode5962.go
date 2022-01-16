/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-08 23:21:04
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-08 23:51:30
 */
package myleetcode

func longestPalindrome(words []string) int {
	ra := make(map[string][]string)

	for _, j := range words {
		r := j
		if j[0] > j[1] {
			r = string([]byte{j[1], j[0]})
		}
		if s, ok := ra[r]; ok {
			s = append(s, j)
			ra[r] = s
		} else {
			ra[r] = []string{j}
		}
	}

	ddcount := 0
	dccountMax := 0
	single := false
	for _, j := range ra {
		if len(j) == 1 {
			if j[0][0] != j[0][1] {
				continue
			}

		}
		m1count := 0
		m2count := 0
		m3count := 0
		for _, m := range j {
			if m[0] < m[1] {
				m1count++
			} else if m[0] == m[1] {
				m3count++
			} else {
				m2count++
			}
		}
		if m3count > 0 {
			if m3count%2 == 0 {
				dccountMax = dccountMax + m3count
			} else {
				single = true
				dccountMax = dccountMax + m3count - 1
			}
		}
		if m1count > m2count {
			ddcount = ddcount + m2count
		} else {
			ddcount = ddcount + m1count
		}
	}
	if single {
		dccountMax++
	}
	return ddcount*2*2 + dccountMax*2
}
