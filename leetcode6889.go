package main

func sumOfSquares(nums []int) int {
	z := len(nums)
	result := 0
	for i := range nums {
		if z%(i+1) == 0 {
			result = result + nums[i]*nums[i]
		}
	}
	return result
}
