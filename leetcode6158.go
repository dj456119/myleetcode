package myleetcode

import "sort"

func shiftingLetters(s string, shifts [][]int) string {
	t := make([]int, len(s))
	// for i := range shifts {
	// 	for start := shifts[i][0]; start <= shifts[i][1]; start++ {
	// 		if shifts[i][2] == 0 {
	// 			t[start]--
	// 		} else {
	// 			t[start]++
	// 		}
	// 	}
	// }
	sort.Slice(shifts, func(i, j int) bool {
		if shifts[i][0] < shifts[j][0] {
			return true
		} else if shifts[i][0] == shifts[j][0] {
			return shifts[i][1] < shifts[j][1]
		}
		return false
	})
	for i := range shifts {
		if i == len(shifts)-1 {
			break
		}
		if shifts[i][0] == shifts[i+1][0] && shifts[1][1] == shifts[i+1][1] {
			if shifts[i][2] == shifts[i][2] {

			}
		}
	}
	f := make([]byte, 26*3)
	start := 'a'
	for i := 0; i <= 25; i++ {
		f[i] = byte(int(start) + i)
	}
	start = 'a'
	for i := 26; i <= 51; i++ {
		f[i] = byte(int(start) + i - 26)
	}
	start = 'a'
	for i := 52; i <= 77; i++ {
		f[i] = byte(int(start) + i - 52)
	}
	// fmt.Println(string(f))
	q := []byte(s)

	for i := range q {
		if t[i] >= 26 || -t[i] >= 26 {
			t[i] = t[i] % 26
		}
		q[i] = f[int(q[i])-int('a')+26+t[i]]
	}
	//   fmt.Println(t, q)
	return string(q)
}
