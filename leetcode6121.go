package myleetcode

import (
	"sort"
	"strconv"
)

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	d := len(nums[0])
	c := make([][][]string, d)
	for i := range c {
		c[i] = make([][]string, len(nums))
		for j := range c[i] {
			c[i][j] = make([]string, 2)
		}
	}
	for i := range nums {
		for j := range nums[i] {
			c[j][i][0] = nums[i][j:]
			c[j][i][1] = strconv.Itoa(i)
		}
	}
	for i := range c {
		sort.Slice(c[i], func(m, n int) bool {
			if c[i][m][0] == c[i][n][0] {
				z1, _ := strconv.Atoi(c[i][m][1])
				z2, _ := strconv.Atoi(c[i][n][1])
				return z1 < z2
			}
			return c[i][m][0] < c[i][n][0]
		})
	}
	//fmt.Println(c)
	result := []int{}
	for i := range queries {
		r := c[len(nums[0])-queries[i][1]][queries[i][0]-1][1]
		zb, _ := strconv.Atoi(r)
		result = append(result, zb)
	}
	return result
}
