package myleetcode

func distinctNames(ideas []string) int64 {
	z := make(map[string][]string)
	for i := range ideas {
		z[ideas[i][1:]] = append(z[ideas[i][1:]], ideas[i])
	}
	result := 0
	var sum1 int64
	for _, v := range z {
		if len(v) > 1 {

		}
	}
}
