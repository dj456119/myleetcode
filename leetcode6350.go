package myleetcode

func maxDivScore(nums []int, divisors []int) int {
	m := make([]int, len(divisors))
	for i := range divisors {
		count := 0
		for j := range nums {
			if nums[j]%divisors[i] == 0 {
				count++
			}
		}
		m[i] = count
	}
	max := 0
	maxValue := 0
	//fmt.Println(m)
	for i := range m {
		if m[i] > maxValue {
			max = i
			maxValue = m[i]
		}
		if m[i] == maxValue {
			if divisors[max] > divisors[i] {
				max = i
			}
		}
	}
	return divisors[max]
}
