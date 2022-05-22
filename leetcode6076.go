package myleetcode

import "sort"

func minimumLines(stockPrices [][]int) int {

	if len(stockPrices) == 1 {
		return 0
	}
	if len(stockPrices) == 2 {
		return 1
	}
	sort.Slice(stockPrices, func(i, j int) bool { return stockPrices[i][0] < stockPrices[j][0] })
	result := 1
	for i := range stockPrices {
		if i == 0 || i == 1 {
			continue
		}
		a2 := stockPrices[i][0]
		a1 := stockPrices[i-1][0]
		a0 := stockPrices[i-2][0]
		b2 := stockPrices[i][1]
		b1 := stockPrices[i-1][1]
		b0 := stockPrices[i-2][1]
		if int64(a2-a1)*int64(b2-b0) != int64(a2-a0)*int64(b2-b1) {
			result++
		}
	}
	return result
}
