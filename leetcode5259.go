package myleetcode

func calculateTax(brackets [][]int, income int) float64 {
	var fc int = income
	var sum float64
	for i := range brackets {
		money := 0
		if i == 0 {
			money = brackets[i][0]
		} else {
			money = brackets[i][0] - brackets[i-1][0]
		}
		if money > fc {
			money = fc
		}
		sum = sum + float64(money)*float64(brackets[i][1])*0.01
		fc = fc - money
		if fc <= 0 {
			break
		}
	}
	return sum
}
