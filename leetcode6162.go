package myleetcode

func garbageCollection(garbage []string, travel []int) int {
	gs := []map[byte]int{}
	for i := range garbage {
		z := make(map[byte]int)
		for j := range garbage[i] {
			z[garbage[i][j]]++
		}
		gs = append(gs, z)
	}
	def := []int{}
	sum := 0
	for i := range travel {
		sum = sum + travel[i]
		def = append(def, sum)
	}
	gcount := 0
	pcount := 0
	mcount := 0
	glast := 0
	plast := 0
	mlast := 0
	for i := range gs {
		gcount = gcount + gs[i]['G']
		mcount = mcount + gs[i]['M']
		pcount = pcount + gs[i]['P']
		if gs[i]['G'] != 0 {
			glast = i
		}
		if gs[i]['M'] != 0 {
			mlast = i
		}
		if gs[i]['P'] != 0 {
			plast = i
		}
	}
	if len(gs) != 1 {
		gcount = gcount + def[glast-1]
		mcount = mcount + def[mlast-1]
		pcount = pcount + def[plast-1]
	}

	result := gcount + mcount + pcount
	return result
}
