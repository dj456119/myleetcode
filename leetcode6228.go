package myleetcode

func twoEditWords(queries []string, dictionary []string) []string {
	result := []string{}
	for i := range queries {
		for j := range dictionary {
			if len(queries[i]) != len(dictionary[j]) {
				continue
			}
			count := 0
			for m := range queries[i] {
				if queries[i][m] != dictionary[j][m] {
					count++
					if count > 2 {
						break
					}
				}
			}
			// fmt.Println(queries[i], dictionary[j], count)
			if count <= 2 {
				result = append(result, queries[i])
				break
			}
		}
	}
	return result
}
