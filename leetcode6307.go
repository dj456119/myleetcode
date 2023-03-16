package main

func passThePillow(n int, time int) int {
	now := 1
	forward := 0
	for time != 0 {
		if forward == 0 {
			now++
		} else {
			now--
		}
		if now == n {
			forward = 1
		} else if now == 1 {
			forward = 0
		}
		time--
	}
	return now
}