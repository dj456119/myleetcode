package myleetcode

func countSpecialNumbers(n int) int {
	weishu := 1
	c := n
	c = c / 10
	for c != 0 {
		weishu++
	}
	sum := 0
	ji := 1
	for i := 1; i <= weishu; i++ {
		ji = ji * i
		sum = sum + ji
	}
	return sum
}
