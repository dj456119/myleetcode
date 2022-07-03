package myleetcode

import (
	"fmt"
	"sort"
)

func countPairs(n int, edges [][]int) int64 {
	aa := make(map[int]bool)
	aaa := make(map[int][]int)
	for i := range edges {
		s := edges[i][0]
		e := edges[i][1]
		if c, ok := aaa[s]; !ok {
			c = make([]int, 0)
			c = append(c, s)
			aaa[s] = c
		} else {
			c = append(c, s)
		}
		if c, ok := aaa[e]; !ok {
			c = make([]int, 0)
			c = append(c, e)
			aaa[e] = c
		} else {
			c = append(c, e)
		}
	}
	fmt.Println(aaa)
	queue := []int{}
	countList := []int{}
	for k, v := range aaa {
		count := 0
		if _, ok := aa[k]; ok {
			continue
		}
		queue = append(queue, v...)
		for len(queue) != 0 {
			z := queue[0]
			if _, ok := aa[z]; ok {
				continue
			}
			count++
			queue = append(queue, aaa[z]...)
		}
		countList = append(countList, count)
	}
	fmt.Println(countList)
	sort.Slice(countList, func(i, j int) bool { return countList[i] < countList[j] })
	if len(countList) == 0 {
		return int64(0)
	}
	var sum int64
	var result int64
	for i := range countList {
		if i == 0 {
			result = int64(countList[0])
			sum = int64(countList[0])
		} else {
			sum = sum * int64(countList[i])
			result = result + sum
		}
	}
	return result
}
