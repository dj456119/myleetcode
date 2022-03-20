package myleetcode

func countCollisions(directions string) int {
	count := 0
	status := directions[0]
	rcount := 0
	for i := range directions {
		if i == 0 {
			if directions[i] == 'R' {
				rcount++
			}
			continue
		}
		if status == 'S' {
			if directions[i] == 'S' {
				continue
			}
			if directions[i] == 'L' {
				status = 'S'
				count++
				continue
			}
			if directions[i] == 'R' {
				status = 'R'
				rcount++
				continue
			}
		}
		if status == 'L' {
			if directions[i] == 'R' {
				rcount++
			}
			status = directions[i]
			continue
		}
		if status == 'R' {
			if directions[i] == 'S' {
				status = 'S'
				count = count + rcount
				rcount = 0
				continue
			}
			if directions[i] == 'R' {
				rcount++
				continue
			}
			if directions[i] == 'L' {

				count = count + rcount
				count = count + 1
				rcount = 0
				status = 'S'
				continue
			}
		}
	}
	return count
}
