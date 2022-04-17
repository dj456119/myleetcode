package myleetcode

func minimumRounds(tasks []int) int {
	taskMap := make(map[int]int)
	for i := range tasks {
		taskMap[tasks[i]]++
	}
	result := 0
	for _, v := range taskMap {
		c := v
		if c == 2 || c == 3 {
			result++
			continue
		}
		if c == 1 {
			return -1
		}
		if c%3 == 0 {
			result = result + c/3
			continue
		}
		if c%3 == 2 {
			result = result + c/3 + 1
			continue
		}
		if c%3 == 1 {
			result = result + c/3 - 1 + 2
			continue
		}
	}
	return result
}
