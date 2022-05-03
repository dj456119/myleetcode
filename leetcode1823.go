package myleetcode

func findTheWinner(n int, k int) int {
	queue := []int{}
	for i := 1; i <= n; i++ {
		queue = append(queue, i)
	}
	for {
		c := k
		for c != 1 {
			m := queue[0]
			queue = queue[1:]
			queue = append(queue, m)
			c--
		}
		queue = queue[1:]
		if len(queue) == 1 {
			return queue[0]
		}
	}
}
