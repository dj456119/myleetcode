package myleetcode

import "math"

func countHousePlacements(n int) int {
	//都放
	var s1 int64
	//都不放
	var s2 int64
	//只放上
	var s3 int64
	//只放下
	var s4 int64

	var sum int64
	for i := 1; i <= n; i++ {
		ns1 := s2
		ns2 := sum
		ns3 := s4 + s2
		ns4 := s3 + s2
		sum = sum + ns1 + ns2 + ns3 + ns4
		s1 = ns1
		s2 = ns2
		s3 = ns3
		s4 = ns4
	}
	return int(sum / (int64(math.Pow10(9)) + 7))
}
