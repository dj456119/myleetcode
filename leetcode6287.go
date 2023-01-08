package myleetcode

func categorizeBox(length int, width int, height int, mass int) string {
	box := []int64{int64(length) * int64(width) * int64(height), int64(length), int64(width), int64(height), int64(mass)}
	z1 := false
	z2 := false
	if box[0] >= 1000000000 || box[1] >= 10000 || box[2] >= 10000 || box[3] >= 10000 || box[4] >= 10000 {
		z1 = true
	}
	if box[4] >= 100 {
		z2 = true
	}
	if z1 && z2 {
		return "Both"
	}
	if z1 {
		return "Bulky"
	}
	if z2 {
		return "Heavy"
	}
	return "Neither"
}
