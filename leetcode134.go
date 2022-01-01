package myleetcode

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	for {
		if start >= len(gas) {
			return -1
		}
		now := 0
		flag := true
		if gas[start] < cost[start] {
			start++
			continue
		}
		for i := start; i < len(gas); i++ {
			if now+gas[i]-cost[i] < 0 {
				start = i
				flag = false
				break
			} else {
				now = now + gas[i] - cost[i]
			}
		}

		if flag {
			for i := 0; i < start; i++ {
				if now+gas[i]-cost[i] < 0 {
					return -1
				} else {
					now = now + gas[i] - cost[i]
				}
			}
			return start
		}
	}
}
