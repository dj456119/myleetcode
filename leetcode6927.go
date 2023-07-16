package main

func minimumIndex(nums []int) int {

	max := 0
	maxNum := 0
	q := make(map[int]int)
	for i := range nums {
		q[nums[i]]++
	}
	for num, count := range q {
		if count > max {
			max = count
			maxNum = num
		}

	}
	left := 0
	right := max
	for i := range nums {
		if i == 0 {
			continue
		}
		if i == len(nums)-1 {
			return -1
		}
		if nums[i] == maxNum {
			left++
			right--
		}
		if left*2 > len(nums[:i+1]) && right*2 > len(nums[i+1:]) {
			return i
		}
	}
	return -1
}
