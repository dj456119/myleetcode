package myleetcode

func minLength(s string) int {
	z := []byte(s)
	haveABorCD := true
	for haveABorCD {
		haveABorCD = false
		z1 := []byte{}
		for i := range z {

			if i < len(z)-1 {
				if z[i] == 'A' && z[i+1] == 'B' {
					haveABorCD = true
					continue
				} else if z[i] == 'C' && z[i+1] == 'D' {
					haveABorCD = true
					continue
				}
			}
			if i != 0 {
				if z[i] == 'B' && z[i-1] == 'A' {
					haveABorCD = true
					continue
				} else if z[i] == 'D' && z[i-1] == 'C' {
					haveABorCD = true
					continue
				}
			}
			z1 = append(z1, z[i])
		}
		z = z1
		// fmt.Println(string(z))
	}
	return len(z)
}
