package myleetcode

import "strconv"

func haveConflict(event1 []string, event2 []string) bool {
	eve1 := [2]int64{}
	eve2 := [2]int64{}
	s1, _ := strconv.ParseInt(event1[0][:2], 10, 32)
	f1, _ := strconv.ParseInt(event1[0][3:], 10, 32)
	eve1[0] = s1*60 + f1
	s11, _ := strconv.ParseInt(event1[1][:2], 10, 32)
	f11, _ := strconv.ParseInt(event1[1][3:], 10, 32)
	eve1[1] = s11*60 + f11

	s2, _ := strconv.ParseInt(event2[0][:2], 10, 32)
	f2, _ := strconv.ParseInt(event2[0][3:], 10, 32)
	eve2[0] = s2*60 + f2
	s21, _ := strconv.ParseInt(event2[1][:2], 10, 32)
	f21, _ := strconv.ParseInt(event2[1][3:], 10, 32)
	eve2[1] = s21*60 + f21
	if eve1[1] < eve2[0] {
		return false
	}
	if eve2[1] < eve1[0] {
		return false
	}
	return true
}
