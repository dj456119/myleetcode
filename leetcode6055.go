package myleetcode

func convertTime(current string, correct string) int {
	time1 := int((current[0]-'0')*10+(current[1]-'0'))*60 + int(current[3]-'0')*10 + int(current[4]-'0')
	t2 := (correct[0]-'0')*10 + (correct[1] - '0')
	time2 := int(t2)*60 + int(correct[3]-'0')*10 + int(correct[4]-'0')
	result := 0
	for time2 != time1 {
		c := time2 - time1
		if c == 60 || c == 5 || c == 15 || c == 1 {
			return result + 1
		}
		if c > 60 {
			time2 = time2 - 60
		}
		if c < 60 && c > 15 {
			time2 = time2 - 15
		}
		if c < 15 && c > 5 {
			time2 = time2 - 5
		}
		if c < 5 && c > 1 {
			time2 = time2 - 1
		}
		result++
	}
	return result
}
